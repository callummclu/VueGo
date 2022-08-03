package helpers

import "go-vue-ecommerce-site/models"

func RemoveIndex(s []models.Product, index int) []models.Product {
	return append(s[:index], s[index+1:]...)
}
