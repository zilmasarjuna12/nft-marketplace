package usecase

import (
	"context"
	"errors"
	"log"
	"nft-marketplace/entity"
	repository_mysql "nft-marketplace/repository/mysql"

	uuid "github.com/satori/go.uuid"
)

type (
	IItemUsecase interface {
		Get(ctx context.Context, query entity.ItemQuery) (items []entity.Item, err error)
		Create(ctx context.Context, creatorId string, input entity.ItemInput) (item *entity.Item, err error)
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

func (usecase *itemUsecase) Create(ctx context.Context, creatorId string, input entity.ItemInput) (item *entity.Item, err error) {
	creatorUuid, _ := uuid.FromString(creatorId)

	item = &entity.Item{
		Name:            *input.Name,
		Rating:          *input.Rating,
		Category:        *input.Category,
		Image:           *input.Image,
		Price:           *input.Price,
		Availibility:    *input.Price,
		ReputationValue: *input.Reputation,
		CreatorID:       creatorUuid,
	}

	item.SetReputationBadge()

	item, err = usecase.itemMysql.Save(ctx, item)

	if err != nil {
		log.Printf("Got usecase.itemMysql.Save Error %v", err)

		return
	}

	return
}
