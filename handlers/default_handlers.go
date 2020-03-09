package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo"
	"github.com/theerfan/AUT-95-Gradsite-BackEnd/graph"
	"github.com/theerfan/AUT-95-Gradsite-BackEnd/graph/generated"
)

func GraphqlHandler() echo.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

}

func PlaygroundHandler() echo.HandlerFunc {

	h := handler.Playground("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

}