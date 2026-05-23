package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"social-media-backend-1/internal/observability"
	"social-media-backend-1/internal/outers/container"
	"social-media-backend-1/internal/outers/deliveries/graphqls"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ravilushqa/otelgqlgen"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	port := os.Getenv("BACKEND_1_PORT")

	// Initialize OpenTelemetry
	shutdown, err := observability.InitOpenTelemetry(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize OpenTelemetry: %v", err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("failed to shutdown OpenTelemetry: %v", err)
		}
	}()

	rootContainer := container.NewRootContainer()
	// Start EDFS subscription listener (Cosmo Router subscriptions)
	go rootContainer.GatewayContainer.EDFGateway.Start(context.Background())

	resolver := graphqls.NewResolver(rootContainer)
	srv := handler.New(graphqls.NewExecutableSchema(graphqls.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{})
	srv.AddTransport(transport.MultipartForm{})

	srv.Use(extension.Introspection{})
	srv.Use(otelgqlgen.Middleware())

	mux := http.NewServeMux()
	mux.Handle("/graphql", rootContainer.MiddlewareContainer.AuthMiddleware.Authenticate(srv))
	mux.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))

	addr := "0.0.0.0:" + port

	// Wrap with otelhttp to extract trace context from Router
	handler := otelhttp.NewHandler(mux, "social-media-backend-1")

	httpServer := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Fatal(httpServer.ListenAndServe())
}
