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
		GetByID(ctx context.Context, itemID string) (item *entity.Item, err error)
		Update(ctx context.Context, itemID string, input entity.ItemUpdate) (item *entity.Item, err error)
		Delete(ctx context.Context, itemID string) (err error)
		Purchased(ctx context.Context, buyerID, itemID string) (err error)
	}

	itemUsecase struct {
		itemMysql repository_mysql.IItemMysql
		userMysql repository_mysql.IUserMysql
	}
)

func NewItemUsecase(itemMysql repository_mysql.IItemMysql, userMysql repository_mysql.IUserMysql) IItemUsecase {
	return &itemUsecase{itemMysql, userMysql}
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

func (usecase *itemUsecase) GetByID(ctx context.Context, itemID string) (item *entity.Item, err error) {
	item, err = usecase.itemMysql.GetByID(ctx, itemID)

	if err != nil {
		log.Printf("Got usecase.itemMysql.GetByID Error %v", err)

		return
	}

	return
}

func (usecase *itemUsecase) Update(ctx context.Context, itemID string, input entity.ItemUpdate) (item *entity.Item, err error) {
	item, err = usecase.itemMysql.Update(ctx, itemID, input)

	if err != nil {
		log.Printf("Got usecase.itemMysql.Save Error %v", err)

		return
	}

	return
}

func (usecase *itemUsecase) Delete(ctx context.Context, itemID string) (err error) {
	err = usecase.itemMysql.Delete(ctx, itemID)

	if err != nil {
		log.Printf("Got usecase.itemMysql.Delete Error %v", err)

		return
	}

	return
}

func (usecase *itemUsecase) Purchased(ctx context.Context, buyerID, itemID string) (err error) {
	_, err = usecase.userMysql.GetByID(ctx, buyerID)

	if err != nil {
		if err.Error() == "record not found" {
			err = errors.New("user not found")
			return
		}

		log.Printf("Got usecase.itemMysql.GetByID Error %v", err)

		return
	}

	item, err := usecase.itemMysql.GetByID(ctx, itemID)

	if err != nil {
		if err.Error() == "record not found" {
			err = errors.New("item not found")
			return
		}

		log.Printf("Got usecase.itemMysql.GetByID Error %v", err)
		return
	}

	if !item.IsAvailable() {
		err = errors.New("empty item")

		return
	}

	err = usecase.itemMysql.Purchased(ctx, buyerID, itemID)

	if err != nil {
		log.Printf("Got usecase.itemMysql.Purchased Error %v", err)

		return
	}

	return
}
