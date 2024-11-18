package storage

import (
	"context"
	"first-app/common"
	"first-app/modules/employee_handle_order/models"
)

func (s *sqlStore) GetListEmployeeHandleOrder(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.GetEmployeeHandleOrder, error){
	var result []models.GetEmployeeHandleOrder

	err := s.db.Order("id asc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Preload("User").
		Find(&result).Error
	
	if err != nil {
		return nil, err
	}

	return result, nil
}
