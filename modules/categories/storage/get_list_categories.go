package storage

import (
	"context"
	"first-app/common"
	"first-app/modules/categories/models"
)

func (s *sqlStore) GetListCategories(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Categories, error) {
	var result []models.Categories

	db := s.db

	if err := s.db.Table(models.Categories{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id asc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
