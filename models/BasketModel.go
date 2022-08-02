package models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	Items []Product `json:"items"`
}

func (Basket) TableName() string {
	return "baskets"
}
