package storage

import (
	"context"
	"first-app/common"
	"first-app/modules/tables/models"
)

func (s *sqlStore) GetListTable(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Table, error) {
	var result []models.Table

	db := s.db

	if err := s.db.Table(models.Table{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id asc").
		Where(map[string]interface{}{"deleted_at": nil}).
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
