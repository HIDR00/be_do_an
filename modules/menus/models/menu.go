package models

import (
	"errors"
	"first-app/common"
)

var (
	ErrTitleIsBlank= errors.New("title cannot be blank")
	ErrMenuDeleted= errors.New("Menu cannot be deleted")
)

type Menu struct {
	common.SQLModel
	Name     string  `json:"name" gorm:"column:name;"`
	Price    float64 `json:"price" gorm:"column:price;"`
	ImageUrl string  `json:"image_url" gorm:"column:image_url;"`
	CategoriesId int `json:"categories_id" gorm:"column:categories_id;"`
	Discription string `json:"discription" gorm:"column:discription;"`
}
