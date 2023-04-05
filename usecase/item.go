package usecase

import (
	"context"
	"errors"
	"log"
	"nft-marketplace/entity"
	repository_mysql "nft-marketplace/repository/mysql"
)

type (
	IItemUsecase interface {
		Get(ctx context.Context, query entity.ItemQuery) (items []entity.Item, err error)
	}

	itemUsecase struct {
		itemMysql repository_mysql.IItemMysql
	}
)

func NewItemUsecase(itemMysql repository_mysql.IItemMysql) IItemUsecase {
	return &itemUsecase{itemMysql}
}

func (usecase *itemUsecase) Get(ctx context.Context, query entity.ItemQuery) (items []entity.Item, err error) {
	items, err = usecase.itemMysql.Get(ctx, query)

	if err != nil {
		log.Printf("Got usecase.itemMysql.Get Error %v", err)

		return
	}

	if len(items) == 0 {
		err = errors.New("not found")
	}

	return
}
