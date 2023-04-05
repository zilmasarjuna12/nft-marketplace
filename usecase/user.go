package usecase

import (
	"context"
	"errors"
	"log"
	"nft-marketplace/entity"
	repository_mysql "nft-marketplace/repository/mysql"
)

type (
	IUserUsecase interface {
		Get(ctx context.Context) (users []entity.User, err error)
	}

	userUsecase struct {
		userMysql repository_mysql.IUserMysql
	}
)

func NewUserUsecase(userMysql repository_mysql.IUserMysql) IUserUsecase {
	return &userUsecase{userMysql}
}

func (usecase *userUsecase) Get(ctx context.Context) (users []entity.User, err error) {
	users, err = usecase.userMysql.Get(ctx)

	if err != nil {
		log.Printf("Got usecase.itemMysql.Get Error %v", err)

		return
	}

	if len(users) == 0 {
		err = errors.New("not found")
	}

	return
}
