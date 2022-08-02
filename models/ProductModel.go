package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func (Product) TableName() string {
	return "products"
}
