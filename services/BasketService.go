package services

import (
	"go-vue-ecommerce-site/configs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var basketCollection *mongo.Collection = configs.GetCollection(configs.DB, "baskets")

func CreateBasket() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var basket models.Basket

		// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&basket).Error; err != nil {
		// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		// 		return
		// 	}

		// 	c.JSON(http.StatusOK, gin.H{"data": basket})
	}
}

func UpdateBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var basket models.Basket
		// if err := models.DB.Where("ID = ?", c.Param("id")).First(&basket).Error; err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		// 	return
		// }

		// var input UpdateBasketInput
		// if err := c.ShouldBindJSON(&input); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// models.DB.Model(&basket).Updates(input)

		// c.JSON(http.StatusOK, gin.H{"data": basket})
	}
}
