package biz

import (
	"context"
	"first-app/modules/users/models"
	"strings"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *models.User) error
}

type CreateUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *CreateUserBiz {
	return &CreateUserBiz{store: store}
}

func (biz *CreateUserBiz) CreateNewUser(ctx context.Context, data *models.User) error {
	email := strings.TrimSpace(data.Email)
	password := strings.TrimSpace(data.Email)
	if email == "" || password == "" {
		return models.ErrTitleIsBlank
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
