package models

import (
	"errors"
	"first-app/common"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrTableDeleted = errors.New("Table cannot be deleted")
)

type Table struct {
	common.SQLModel
	Status int `json:"status" gorm:"column:status;"`
}
