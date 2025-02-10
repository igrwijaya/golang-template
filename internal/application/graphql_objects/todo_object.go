package graphqlobjects

import (
	"igrwijaya-go-template/internal/domain/todo"

	"github.com/graphql-go/graphql"
)

func TodoObjectGraph() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Todo",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						todo := params.Source.(todo.Todo)

						return todo.Id, nil
					},
				},
				"Title": &graphql.Field{
					Type: graphql.String,
				},
				"Description": &graphql.Field{
					Type: graphql.String,
				},
				"CreatedAt": &graphql.Field{
					Type: graphql.DateTime,
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						todo := params.Source.(todo.Todo)

						return todo.CreatedAt, nil
					},
				},
				"UpdatedAt": &graphql.Field{
					Type: graphql.DateTime,
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						todo := params.Source.(todo.Todo)

						return todo.CreatedAt, nil
					},
				},
			},
		},
	)
}
