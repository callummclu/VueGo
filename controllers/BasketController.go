package controllers

import (
	"go-vue-ecommerce-site/middleware"
	"go-vue-ecommerce-site/models"

	"github.com/gin-gonic/gin"
)

func BasketController() {
	api := Router.Group("basket")
	{
		api.Use(middleware.CORSMiddleware())

		api.POST("", func(c *gin.Context) {
			var basket models.Basket
			c.BindJSON(&basket)
			result := models.DB.Create(&basket)
			c.JSON(200, gin.H{
				"checkout": result,
			})
		})
	}
}
