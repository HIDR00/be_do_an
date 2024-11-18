package storage

import (
	"context"
	"first-app/modules/users/models"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *models.User) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}