package models

import (
	"errors"
	"first-app/common"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrOrderDeleted = errors.New("Order cannot be deleted")
)

type Table struct {
	Id      int `json:"id" gorm:"primaryKey;column:id"`
	Status  int `json:"status" gorm:"column:status;"`
	PayType int `json:"pay_type" gorm:"column:pay_type;"`
}

type TableCreate struct {
	Id      int           `json:"id" gorm:"primaryKey;column:id"`
	Status  int           `json:"status" gorm:"column:status;"`
	PayType int           `json:"pay_type" gorm:"column:pay_type;"`
	Orders  []OrderCreate `json:"orders" gorm:"foreignKey:TableID"`
}

type OrderCreate struct {
	common.SQLModel
	TableID  uint `json:"table_id" gorm:"column:table_id"`
	MenuID   uint `json:"menu_id" gorm:"column:menu_id"`
	Quantity int  `json:"quantity" gorm:"column:quantity;"`
}

type TableGet struct {
	Id      int        `json:"id" gorm:"primaryKey;column:id"`
	Status  int        `json:"status" gorm:"column:status;"`
	PayType int        `json:"pay_type" gorm:"column:pay_type;"`
	Orders  []OrderGet `json:"orders" gorm:"foreignKey:TableID"`
}

type OrderGet struct {
	common.SQLModel
	TableID  uint    `json:"-" gorm:"column:table_id"`
	MenuID   uint    `json:"-" gorm:"column:menu_id"`
	Quantity int     `json:"quantity" gorm:"column:quantity;"`
	Menu     MenuGet `json:"menu" gorm:"foreignKey:MenuID"`
}

type MenuGet struct {
	common.SQLModel
	Name     string  `json:"name" gorm:"column:name;"`
	Price    float64 `json:"price" gorm:"column:price;"`
	ImageUrl string  `json:"image_url" gorm:"column:image_url;"`
}

type TotalOrder struct {
	TotalRevenueByDay   float64  `json:"totalRevenueByDay"`
	TotalRevenueByMonth float64  `json:"totalRevenueByMonth"`
	TotelOrderByDay     int `json:"totalOrderByDay"`
	TotelOrderByMonth   int  `json:"TotelOrderByMonth"`
}
