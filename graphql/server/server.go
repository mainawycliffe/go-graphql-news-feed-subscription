package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/websocket"
	"github.com/mainawycliffe/go-graphql-news-feed-subscription/graphql"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			// being lazy, allow all origins
			return true
		},
		AllowCredentials: true,
	})

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/", c.Handler(handler.Playground("GraphQL playground", "/query")))
	http.Handle("/query", handler.GraphQL(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: graphql.GraphQLServer(),
			},
		),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", getPort())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil))
}
