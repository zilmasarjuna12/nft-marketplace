package graph

import (
	"log"
	"nft-marketplace/config"
	repository_mysql "nft-marketplace/repository/mysql"

	"nft-marketplace/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	itemUsecase usecase.IItemUsecase
}

func NewResolver() Config {
	db, err := config.NewDatabase()

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	itemPostgres := repository_mysql.NewItemMysql(db)

	itemUsecase := usecase.NewItemUsecase(itemPostgres)

	r := Resolver{
		itemUsecase: usecase.NewItemUsecase(itemUsecase),
	}

	return Config{
		Resolvers: &r,
	}
}
