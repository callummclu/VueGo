package models

type Product struct {
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func (Product) TableName() string {
	return "products"
}
