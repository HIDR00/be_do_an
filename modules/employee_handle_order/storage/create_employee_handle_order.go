package storage

import (
	"context"
	"first-app/modules/employee_handle_order/models"
)

func (s *sqlStore) CreateEmployeeHandleOrder(ctx context.Context, data *models.EmployeeHandleOrder) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}