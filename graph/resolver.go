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
	userUsecase usecase.IUserUsecase
}

func NewResolver() Config {
	db, err := config.NewDatabase()

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	itemMysql := repository_mysql.NewItemMysql(db)
	userMysql := repository_mysql.NewUserMysql(db)

	itemUsecase := usecase.NewItemUsecase(itemMysql, userMysql)
	userUsecase := usecase.NewUserUsecase(userMysql)

	r := Resolver{
		itemUsecase: itemUsecase,
		userUsecase: userUsecase,
	}

	return Config{
		Resolvers: &r,
	}
}
