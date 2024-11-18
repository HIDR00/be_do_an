package biz

import (
	"context"
	"first-app/modules/users/models"
)

type UpdateUserByIdStorage interface {
	GetUserById(ctx context.Context,cond map[string]interface{}) (*models.User,error)
	UpdateUserById(ctx context.Context,cond map[string]interface{},dataUpdate *models.User) error
}

type UpdateUserByIdBiz struct {
	store UpdateUserByIdStorage
}

func NewUpdateUserByIdBiz(store UpdateUserByIdStorage) *UpdateUserByIdBiz {
	return &UpdateUserByIdBiz{store: store}
}

func (biz *UpdateUserByIdBiz) UpdateNewUserById(ctx context.Context,id int,dataUpdate *models.User) error {

	data, err := biz.store.GetUserById(ctx,map[string]interface{}{"id":id}); 

	if err != nil {
		return err
	}

	if data.Email == "" {
		return models.ErrItemDeleted
	}

	if err := biz.store.UpdateUserById(ctx,map[string]interface{}{"id": id},dataUpdate); err != nil {
		return err
	}
	return nil
}
