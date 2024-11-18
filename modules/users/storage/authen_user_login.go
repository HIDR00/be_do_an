package storage

import (
	"context"
	"first-app/modules/users/models"
)

func (s *sqlStore) AuthenUser(ctx context.Context,cond map[string]interface{}) (*models.User,error) {
	var data models.User

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil,err
	}

	return &data,nil
}