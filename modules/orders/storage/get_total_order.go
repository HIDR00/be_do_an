package storage

import (
	"context"
	"first-app/modules/orders/models"
)

func (s *sqlStore) GetTotalOrder(ctx context.Context) ([]models.OrderGet, error) {
	var data []models.OrderGet

	if err := s.db.Preload("Menu").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
