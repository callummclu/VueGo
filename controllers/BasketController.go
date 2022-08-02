package controllers

import (
	"go-vue-ecommerce-site/middleware"
	_ "go-vue-ecommerce-site/models"
	"go-vue-ecommerce-site/services"
)

func BasketController() {
	api := Router.Group("basket")
	{
		api.Use(middleware.CORSMiddleware())

		api.POST("", services.CreateBasket())
		api.GET(":id", services.GetBasket())
		api.POST(":id", services.UpdateBasket())
	}
}
