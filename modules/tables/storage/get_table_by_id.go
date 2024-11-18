package storage

import (
	"context"
	"first-app/modules/tables/models"
)

func (s *sqlStore) GetTableById(ctx context.Context,cond map[string]interface{}) (*models.Table,error) {
	var data models.Table

	if err := s.db.Where(cond).
		First(&data).Error; err != nil {
		return nil,err
	}
	return &data,nil
}