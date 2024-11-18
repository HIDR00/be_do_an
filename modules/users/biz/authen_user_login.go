package biz


import (
	"context"
	"first-app/modules/users/models"
)

type AuthenUserStorage interface {
	AuthenUser(ctx context.Context,cond map[string]interface{}) (*models.User,error)
}

type AuthenUserBiz struct {
	store AuthenUserStorage
}

func NewAuthenUserBiz(store AuthenUserStorage) *AuthenUserBiz {
	return &AuthenUserBiz{store: store}
}

func (biz *AuthenUserBiz) AuthenNewUser(ctx context.Context,data *models.User) (*models.User,error) {

	data,err := biz.store.AuthenUser(ctx,map[string]interface{}{"email":data.Email,"password":data.Password}); 
	if err != nil {
		return nil,err
	}

	return data,nil 
}
