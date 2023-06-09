package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nft-marketplace/graph"
	"nft-marketplace/migrations"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	migrations.Migration()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.NewResolver()))
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		rc := graphql.GetFieldContext(ctx)
		fmt.Println("Entered", rc.Object, rc.Field.Name)
		res, err = next(ctx)
		fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
		return res, err
	})

	http.Handle("/", playground.Handler("Nft", "/query"))
	http.Handle("/query", srv)

	log.Println("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
