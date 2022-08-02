package models

type Basket struct {
	Id    int       `json:"id" gorm:"PrimaryKey"`
	Items []Product `json:"items"`
}
