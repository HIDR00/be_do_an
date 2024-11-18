package biz


import (
	"context"
	"first-app/modules/users/models"
)

type DeleteUserByIdStorage interface {
	GetUserById(ctx context.Context,cond map[string]interface{}) (*models.User,error)
	DeleteUserById(ctx context.Context,cond map[string]interface{}) error
}

type DeleteUserByIdBiz struct {
	store DeleteUserByIdStorage
}

func NewDeleteUserByIdBiz(store DeleteUserByIdStorage) *DeleteUserByIdBiz {
	return &DeleteUserByIdBiz{store: store}
}

func (biz *DeleteUserByIdBiz) DeleteNewUserById(ctx context.Context,id int) error {

	data, err := biz.store.GetUserById(ctx,map[string]interface{}{"id":id}); 

	if err != nil {
		return err
	}

	if data.Email == "" {
		return models.ErrItemDeleted
	}

	if err := biz.store.DeleteUserById(ctx,map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
