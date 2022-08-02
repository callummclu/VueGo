package controllers

import (
	"go-vue-ecommerce-site/middleware"

	"github.com/gin-gonic/gin"
)

func CheckoutController() {
	api := Router.Group("checkout")
	{
		api.Use(middleware.CORSMiddleware())

		api.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"checkout": "empty",
			})
		})
	}

	/*
		CHECKOUT ITEMS IN BASKET
			- CHECK BASKET IS VALID
			- DELETE BASKET
	*/
}
