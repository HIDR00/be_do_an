package storage

import (
	"context"
	"first-app/modules/orders/models"
)

func (s *sqlStore) CreateOrder(ctx context.Context, data *[]models.OrderCreate,cond map[string]interface{}) error {

	db := s.db
    if err := db.Create(&data).Error ; err != nil {
		return err
	}

	return nil
}