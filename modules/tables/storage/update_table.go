package storage

import (
	"context"
	"first-app/modules/tables/models"
)

func (s *sqlStore) UpdateTableById(ctx context.Context,cond map[string]interface{},dataUpdate *models.Table) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}