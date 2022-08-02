package controllers

import (
	"go-vue-ecommerce-site/middleware"
	"go-vue-ecommerce-site/services"
)

func CheckoutController() {
	api := Router.Group("checkout")
	{
		api.Use(middleware.CORSMiddleware())

		api.POST(":id", services.CheckoutBasket())
	}
}
