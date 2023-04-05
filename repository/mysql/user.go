package repository_mysql

import (
	"context"
	"log"
	"nft-marketplace/entity"

	"gorm.io/gorm"
)

type (
	IUserMysql interface {
		Get(ctx context.Context) (user []entity.User, err error)
		GetByID(ctx context.Context, userID string) (user *entity.User, err error)
	}

	userMysql struct {
		DB *gorm.DB
	}
)

func NewUserMysql(DB *gorm.DB) IUserMysql {
	return &userMysql{DB}
}

func (repo *userMysql) Get(ctx context.Context) (user []entity.User, err error) {
	db := repo.DB.Debug().WithContext(ctx)

	if err = db.Find(&user).Error; err != nil {
		log.Printf("Failed Find With Error : %v", err)
		return
	}

	return
}

func (repo *userMysql) GetByID(ctx context.Context, userID string) (user *entity.User, err error) {
	db := repo.DB.Debug().WithContext(ctx)

	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		log.Printf("Failed First With Error : %v", err)
		return nil, err
	}

	return
}
