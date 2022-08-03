package models

type Product struct {
	ID          string  `json:"_id", omitempty`
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func (Product) TableName() string {
	return "products"
}
