package controllers

import (
	"go-vue-ecommerce-site/middleware"
	"go-vue-ecommerce-site/services"
)

func BasketController() {
	api := Router.Group("basket")
	{
		api.Use(middleware.CORSMiddleware())

		api.POST("", services.CreateBasket())
		api.GET(":id", services.GetBasket())
		api.POST(":id", services.AddToBasket())
		api.PUT(":id", services.UpdateBasket())
		api.DELETE(":id", services.DeleteOneFromBasket())
	}
}
