package biz

import (
	"context"
	"first-app/modules/users/models"
)

type GetUserByIdStorage interface {
	GetUserById(ctx context.Context,cond map[string]interface{}) (*models.User,error)
}

type GetUserByIdBiz struct {
	store GetUserByIdStorage
}

func NewGetUserByIdBiz(store GetUserByIdStorage) *GetUserByIdBiz {
	return &GetUserByIdBiz{store: store}
}

func (biz *GetUserByIdBiz) GetNewUserById(ctx context.Context,id int) (*models.User,error) {

	data, err := biz.store.GetUserById(ctx,map[string]interface{}{"id":id}); 
	if err != nil {
		return nil,err
	}

	return data,nil
}
