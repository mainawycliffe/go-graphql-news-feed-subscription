package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	badger "github.com/dgraph-io/badger"
	"github.com/mainawycliffe/go-graphql-news-feed-subscription/graphql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: graphql.GetResolver(db),
			},
		),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
