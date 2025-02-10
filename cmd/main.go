package main

import (
	"igrwijaya-go-template/internal/application/graphqlschemas"
	"igrwijaya-go-template/internal/infrastructure/db"
	"igrwijaya-go-template/internal/infrastructure/migrations"
	"igrwijaya-go-template/internal/infrastructure/repositories"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	migrations.MigrateEntity()

	appDb := db.NewAppDb()
	todoRepo := repositories.NewTodoRepository(appDb)
	todoGraphql := graphqlschemas.NewTodoGraphql(todoRepo)

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    todoGraphql.Query(),
			Mutation: todoGraphql.Mutation(),
		},
	)

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
