package repository_mysql

import (
	"context"
	"errors"
	"log"
	"nft-marketplace/entity"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type (
	IItemMysql interface {
		Get(ctx context.Context, query entity.ItemQuery) (item []entity.Item, err error)
		Save(ctx context.Context, item *entity.Item) (*entity.Item, error)
		Update(ctx context.Context, itemID string, input entity.ItemUpdate) (*entity.Item, error)
		GetByID(ctx context.Context, itemID string) (item *entity.Item, err error)
		Delete(ctx context.Context, itemID string) (err error)
		Purchased(ctx context.Context, buyerID, itemID string) (err error)
	}

	itemMysql struct {
		DB *gorm.DB
	}
)

func NewItemMysql(DB *gorm.DB) IItemMysql {
	return &itemMysql{DB}
}

func (repo *itemMysql) Get(ctx context.Context, query entity.ItemQuery) (items []entity.Item, err error) {
	db := repo.DB.Debug().WithContext(ctx)

	if query.Category != nil {
		db = db.Where("category = ?", query.Category)
	}

	if err = db.Preload("Creator").Find(&items).Error; err != nil {
		log.Printf("Failed Find With Error : %v", err)
		return
	}

	return
}

func (repo *itemMysql) Save(ctx context.Context, item *entity.Item) (*entity.Item, error) {
	db := repo.DB.Debug().WithContext(ctx)
	if err := db.Save(&item).Error; err != nil {
		log.Printf("Failed Save With Error : %v", err)
		return nil, err
	}

	return item, nil
}

func (repo *itemMysql) Update(ctx context.Context, itemID string, input entity.ItemUpdate) (item *entity.Item, err error) {
	db := repo.DB.Debug().WithContext(ctx)

	if err := db.Where("id = ?", itemID).First(&item).Error; err != nil {
		log.Printf("Failed First With Error : %v", err)
		return nil, err
	}

	// edit item
	item.Edit(input)

	if err := db.Save(&item).Error; err != nil {
		log.Printf("Failed Save With Error : %v", err)
		return nil, err
	}

	return item, nil
}

func (repo *itemMysql) GetByID(ctx context.Context, id string) (item *entity.Item, err error) {
	db := repo.DB.Debug().WithContext(ctx)

	if err := db.Where("id = ?", id).Preload("Creator").First(&item).Error; err != nil {
		log.Printf("Failed First With Error : %v", err)
		return nil, err
	}

	return
}

func (repo *itemMysql) Delete(ctx context.Context, itemID string) (err error) {
	db := repo.DB.Debug().WithContext(ctx)

	var transaction []entity.Transaction

	if err = db.Where("item_id = ?", itemID).Find(&transaction).Error; err != nil {
		log.Printf("Failed Delete With Error : %v", err)

		return
	}

	if len(transaction) > 0 {
		err = errors.New("not acceptable")
		return
	}

	if err = db.Where("id = ?", itemID).Delete(&entity.Item{}).Error; err != nil {
		log.Printf("Failed Delete With Error : %v", err)
		return
	}

	return
}

func (repo *itemMysql) Purchased(ctx context.Context, buyerID, itemID string) (err error) {
	tx := repo.DB.Debug().WithContext(ctx).Begin()

	var (
		item        *entity.Item
		transaction *entity.Transaction
	)

	// save transaction
	buyerUuid, _ := uuid.FromString(buyerID)
	itemUuid, _ := uuid.FromString(itemID)

	transaction = &entity.Transaction{
		BuyerID: buyerUuid,
		ItemID:  itemUuid,
	}

	if err = tx.Save(&transaction).Error; err != nil {
		tx.Rollback()

		log.Printf("Failed Transaction Save With Error : %v", err)
		return
	}

	// update availibility
	if err := tx.Where("id = ?", itemID).First(&item).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed First With Error : %v", err)
		return err
	}

	// decrement
	item.DecrementAvailibility()

	if err := tx.Save(&item).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed Save With Error : %v", err)
		return err
	}

	tx.Commit()

	return
}
