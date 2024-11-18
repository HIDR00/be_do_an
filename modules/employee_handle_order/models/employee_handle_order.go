package models

import (
	"errors"
	"first-app/common"
)

var (
	ErrTitleIsBlank = errors.New("email or password cannot be blank")
	ErrItemDeleted  = errors.New("title cannot be deleted")
)

type EmployeeHandleOrder struct {
	common.SQLModel
	TotalPrice  float64 `json:"total_price" gorm:"column:total_price;"`
	TableId     int     `json:"table_id" gorm:"column:table_id;"`
	TableStatus int     `json:"table_status" gorm:"column:table_status;"`
	UserId      int     `json:"user_id" gorm:"column:user_id;"`
}

type GetEmployeeHandleOrder struct {
	common.SQLModel
	TotalPrice  float64 `json:"total_price" gorm:"column:total_price;"`
	TableId     int     `json:"table_id" gorm:"column:table_id;"`
	TableStatus int     `json:"table_status" gorm:"column:table_status;"`
	UserId      int     `json:"-" gorm:"column:user_id;"`
	User        User    `json:"user" gorm:"foreignKey:UserId;references:Id"`
}

type User struct {
	common.SQLModel
	Name  string `json:"name" gorm:"column:name;"`
	Email string `json:"email" gorm:"c olumn:email;"`
}
