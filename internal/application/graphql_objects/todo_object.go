package graphqlobjects

import "github.com/graphql-go/graphql"

func TodoObjectGraph() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Todo",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Title": &graphql.Field{
					Type: graphql.String,
				},
				"Description": &graphql.Field{
					Type: graphql.String,
				},
				"CreatedAt": &graphql.Field{
					Type: graphql.DateTime,
				},
				"UpdatedAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
