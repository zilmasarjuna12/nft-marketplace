package repository_mysql

import (
	"context"
	"log"
	"nft-marketplace/entity"

	"gorm.io/gorm"
)

type (
	IItemMysql interface {
		Get(ctx context.Context, query entity.ItemQuery) (item []entity.Item, err error)
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
