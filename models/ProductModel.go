package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Id          int     `json:"id" gorm:"primaryKey"`
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func (Product) TableName() string {
	return "products"
}
