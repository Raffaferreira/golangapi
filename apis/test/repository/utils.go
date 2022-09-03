package utils

import (
	main "main/apis/test"
	st "main/apis/test/models"
)

func findGrocery(groceries string) (*st.Grocery, error) {
	var grocery st.Grocery
	if result := main.Db.Where(&st.Grocery{Name: groceries}).First(&grocery); result.Error != nil {
		return nil, result.Error
	}
	return &grocery, nil
}

func findQuantity(quantity int) ([]st.Grocery, error) {
	var groceries []st.Grocery
	if result := main.Db.Where("Quantity > ?", quantity).Find(&groceries); result.Error != nil {
		return nil, result.Error
	}
	return groceries, nil
}

func insertGrocery(grocery *st.Grocery) error {
	if result := main.Db.Create(grocery); result.Error != nil {
		return result.Error
	}
	return nil
}
