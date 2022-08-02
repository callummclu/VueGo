package services

import (
	"go-vue-ecommerce-site/configs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id := c.Param("id")
	}
}

func CreateOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var product models.Product
		// c.BindJSON(&product)
		// result := models.DB.Create(&product)
		// c.JSON(200, gin.H{
		// 	"message": result.RowsAffected,
		// })
	}
}

func DeleteOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// result := models.DB.Delete(&models.Product{})

		// c.JSON(200, gin.H{
		// 	"message": result.RowsAffected,
		// })
	}
}

func EditOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
