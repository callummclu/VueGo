package models

type Basket struct {
	Id    string    `json:"id"`
	Items []Product `json:"items"`
	//Total float32   `json:"total"`
}
