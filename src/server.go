package main

import (
	"docker-go/db"
	"docker-go/graph"
	"docker-go/infrastructure"
	"docker-go/usecase"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)



func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Println("failed to connect db")
	}

	ui := infrastructure.NewUserInfrasstructure(db)
	uu := usecase.NewUserUsecase(ui)

	ti := infrastructure.NewTodoInfrasstructure(db)
	tu := usecase.NewTodoUsecase(ti)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserUsacase: uu, TodoUsecase: tu}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
})
	http.Handle("/query", c.Handler(srv))

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
