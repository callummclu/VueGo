package controllers

import (
	"go-vue-ecommerce-site/models"
	"math/rand"

	"github.com/gin-gonic/gin"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func BasketController() {
	api := Router.Group("basket")
	{
		api.POST("", func(c *gin.Context) {
			var basket models.Basket

			result := models.DB.Create(&basket)

			c.JSON(200, gin.H{
				"basket": result.Value,
			})
		})

		api.GET(":id", func(c *gin.Context) {
			id := c.Params.ByName("id")
			c.JSON(200, gin.H{
				"basket": models.DB.First(&models.Basket{}, id),
			})
		})

	}
}
