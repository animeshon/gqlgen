package main

import (
	"log"
	"net/http"

	"github.com/animeshon/gqlgen/example/selection"
	"github.com/animeshon/gqlgen/graphql/handler"
	"github.com/animeshon/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
