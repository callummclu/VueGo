package models

type Order struct {
	ID            string `json:"_id"`
	BasketContent Basket `json:"basket"`
	Status        string `json:"status`
}

func (Order) TableName() string {
	return "orders"
}

