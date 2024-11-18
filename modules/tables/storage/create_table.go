package storage

import (
	"context"
	"first-app/modules/tables/models"
)

func (s *sqlStore) CreateTable(ctx context.Context, data *models.Table) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}