package models

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"title"`
	Price       float32 `json:"price"`
}
