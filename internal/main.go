package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"social-media-backend-1/internal/outers/container"
	"social-media-backend-1/internal/outers/deliveries/graphqls"
)

func main() {
	port := os.Getenv("BACKEND_1_PORT")

	rootContainer := container.NewRootContainer()

	resolver := graphqls.NewResolver(rootContainer)
	srv := handler.New(graphqls.NewExecutableSchema(graphqls.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{})
	srv.AddTransport(transport.MultipartForm{})

	srv.Use(extension.Introspection{})

	mux := http.NewServeMux()
	mux.Handle("/graphql", srv)
	mux.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))

	addr := "0.0.0.0:" + port
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Fatal(httpServer.ListenAndServe())
}
