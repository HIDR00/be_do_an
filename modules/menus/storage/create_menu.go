package storage

import (
	"context"
	"first-app/modules/menus/models"
)

func (s *sqlStore) CreateMenu(ctx context.Context, data *models.Menu) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}