package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewHandler(uh UserHandler, ch ContentHandler) *handler.Handler {
	schema, _ := NewSchema(uh, ch)
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return h
}

func NewSchema(uh UserHandler, ch ContentHandler) (graphql.Schema, error) {
	queryFields := graphql.Fields{
		"user": &graphql.Field{
			Type:    userType,
			Resolve: uh.Query,
		},
		"contents": &graphql.Field{
			Type: graphql.NewList(contentType),
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: ch.Query,
		},
	}
	// mutationFields := graphql.Fields{
	// 	"updateuser": &graphql.Field{
	// 		// Type: graphql.String,
	// 		Type: userType,
	// 		Args: graphql.FieldConfigArgument{
	// 			"name": &graphql.ArgumentConfig{
	// 				Type: graphql.String,
	// 			},
	// 			"nick_name": &graphql.ArgumentConfig{
	// 				Type: graphql.String,
	// 			},
	// 			"desc": &graphql.ArgumentConfig{
	// 				Type: graphql.String,
	// 			},
	// 			"email": &graphql.ArgumentConfig{
	// 				Type: graphql.String,
	// 			},
	// 		},
	// 		Resolve: uh.Mutation,
	// 	},
	// }
	rootQuery := graphql.ObjectConfig{Name: "query", Fields: queryFields}
	// rootMutation := graphql.ObjectConfig{Name: "mutation", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		// Mutation: graphql.NewObject(rootMutation),
	}
	return graphql.NewSchema(schemaConfig)
}
