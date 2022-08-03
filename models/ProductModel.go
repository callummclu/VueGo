package models

type Product struct {
	ID          string  `json:"_id", omitempty`
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Image       string  `json:"string"`
}

func (Product) TableName() string {
	return "products"
}
