package models

type Order struct {
	ID            string `json:"_id"`
	BasketContent Basket `json:"basket"`
	Status        string `json:"status`
}

func (Order) TableName() string {
	return "orders"
}

// Status should either be "PAID" or "SHIPPED"
// Order should be created when checkout is called with a status of "PAID"
