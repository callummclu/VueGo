package models

type Basket struct {
	Items []Product `json:"items"`
}

func (Basket) TableName() string {
	return "baskets"
}
