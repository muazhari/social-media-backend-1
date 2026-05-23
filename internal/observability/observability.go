package observability

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"

	_ "net/http/pprof" // Standard profiling endpoint for Go
)

// InitOpenTelemetry initializes logs, metrics, traces, and profiling
func InitOpenTelemetry(ctx context.Context) (func(context.Context) error, error) {
	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "social-media-backend-1"
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
		resource.WithTelemetrySDK(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithHost(),
	)
	if err != nil {
		return nil, err
	}

	var shutdownFuncs []func(context.Context) error
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			if e := fn(ctx); e != nil {
				err = e
			}
		}
		return err
	}

	// Propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Trace Provider
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return shutdown, err
	}
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

	// Metric Provider
	metricExporter, err := otlpmetricgrpc.New(ctx)
	if err != nil {
		return shutdown, err
	}
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)

	// Log Provider
	logExporter, err := otlploggrpc.New(ctx)
	if err != nil {
		return shutdown, err
	}
	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter)),
		sdklog.WithResource(res),
	)
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)

	// Set as global logger for slog
	globalLogger := otelslog.NewLogger(serviceName, otelslog.WithLoggerProvider(loggerProvider))

	// Use global logger
	slog.SetDefault(globalLogger)

	go func() {
		log.Println("Starting Profiling (pprof) on :6060")
		if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
			log.Printf("Pprof server failed: %v", err)
		}
	}()

	return shutdown, nil
}
