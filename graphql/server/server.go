package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/mainawycliffe/go-graphql-news-feed-subscription/graphql"
)

const defaultPort = "8080"

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	return port
}

func main() {

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: graphql.GraphQLServer(),
			},
		),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", getPort())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil))
}
