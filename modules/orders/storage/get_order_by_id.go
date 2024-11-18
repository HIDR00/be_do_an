package storage

import (
	"context"
	"first-app/modules/orders/models"
)

func (s *sqlStore) GetOrderById(ctx context.Context, cond map[string]interface{}) (*models.TableGet, error) {
	var data models.TableGet

	if err := s.db.Preload("Orders",map[string]interface{}{"deleted_at": nil}).Preload("Orders.Menu").Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
