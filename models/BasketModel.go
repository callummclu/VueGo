package models

type Basket struct {
	Items []Product `json:"items"`
	Total float32   `json:"total"`
}

func (Basket) TableName() string {
	return "baskets"
}
