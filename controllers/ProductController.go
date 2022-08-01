package controllers

import (
	"go-vue-ecommerce-site/models"

	"github.com/gin-gonic/gin"
)

func ProductContoller() {
	api := Router.Group("product")
	{
		api.GET("", func(c *gin.Context) {
			var products []models.Product

			c.JSON(200, gin.H{
				"product": models.DB.Find(&products),
			})
		})

		api.GET(":id", func(c *gin.Context) {
			id := c.Params.ByName("id")

			c.JSON(200, gin.H{
				"product": models.DB.Find(&models.Product{}, id),
			})
		})

		api.POST("", func(c *gin.Context) {
			var product models.Product
			c.BindJSON(&product)
			result := models.DB.Create(&product)
			c.JSON(200, gin.H{
				"message": result.RowsAffected,
			})
		})

		api.DELETE(":id", func(c *gin.Context) {

			id := c.Params.ByName("id")

			result := models.DB.Delete(&models.Product{}, id)

			c.JSON(200, gin.H{
				"message": result.RowsAffected,
			})
		})
	}
}
