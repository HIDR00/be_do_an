package storage

import (
	"context"
	"first-app/common"
	"first-app/modules/menus/models"
)

func (s *sqlStore) GetListMenu(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Menu, error) {
	var result []models.Menu

	db := s.db

	if err := s.db.Table(models.Menu{}.TableName()).Count(&paging.Total).Error; err != nil {
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
