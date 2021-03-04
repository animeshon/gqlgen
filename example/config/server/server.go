package main

import (
	"log"
	"net/http"

	todo "github.com/animeshon/gqlgen/example/config"
	"github.com/animeshon/gqlgen/graphql/handler"
	"github.com/animeshon/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
