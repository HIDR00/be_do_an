package storage

import (
	"context"
	"first-app/modules/users/models"
)

func (s *sqlStore) UpdateUserById(ctx context.Context,cond map[string]interface{},dataUpdate *models.User) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}