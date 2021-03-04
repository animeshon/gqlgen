package main

import (
	"log"
	"net/http"

	"github.com/animeshon/gqlgen/example/dataloader"
	"github.com/animeshon/gqlgen/graphql/handler"
	"github.com/animeshon/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Use(dataloader.LoaderMiddleware)

	router.Handle("/", playground.Handler("Dataloader", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		dataloader.NewExecutableSchema(dataloader.Config{Resolvers: &dataloader.Resolver{}}),
	))

	log.Println("connect to http://localhost:8082/ for graphql playground")
	log.Fatal(http.ListenAndServe(":8082", router))
}
