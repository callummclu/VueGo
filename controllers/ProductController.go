package controllers

import (
	"go-vue-ecommerce-site/models"

	"github.com/gin-gonic/gin"
)

func ProductContoller() {
	api := Router.Group("product")
	{
		// GET ALL PRODUCTS
		api.GET("", func(c *gin.Context) {
			var products []models.Product

			c.JSON(200, gin.H{
				"product": models.DB.Find(&products),
			})
		})

		// GET ONE PRODUCT
		api.GET(":id", func(c *gin.Context) {
			id := c.Params.ByName("id")

			c.JSON(200, gin.H{
				"product": models.DB.Find(&models.Product{}, id),
			})
		})

		// CREATE ONE PRODUCT
		api.POST("", func(c *gin.Context) {
			var product models.Product
			c.BindJSON(&product)
			result := models.DB.Create(&product)
			c.JSON(200, gin.H{
				"message": result.RowsAffected,
			})
		})

		// DELETE ONE PRODUCTS
		api.DELETE(":id", func(c *gin.Context) {

			id := c.Params.ByName("id")

			result := models.DB.Delete(&models.Product{}, id)

			c.JSON(200, gin.H{
				"message": result.RowsAffected,
			})
		})

		// DELETE ALL PRODUCTS

		// EDIT ONE PRODUCT
	}
}
