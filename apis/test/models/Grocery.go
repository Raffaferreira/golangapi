package grocery

import "gorm.io/gorm"

var Db *gorm.DB

type Grocery struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var GroceriesData = []Grocery{
	{Name: "Almod Milk", Quantity: 2},
	{Name: "Apple", Quantity: 6},
}
