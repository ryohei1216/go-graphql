package main

import (
	"docker-go/db"
	"docker-go/graph"
	"docker-go/infrastructure"
	"docker-go/usecase"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Println("failed to connect db")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ui := infrastructure.NewUserInfrasstructure(db)
	uu := usecase.NewUserUsecase(ui)

	ti := infrastructure.NewTodoInfrasstructure(db)
	tu := usecase.NewTodoUsecase(ti)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserUsacase: uu, TodoUsecase: tu}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
