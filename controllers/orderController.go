package controllers

import (
	"go-vue-ecommerce-site/middleware"
	"go-vue-ecommerce-site/services"
)

func OrderController() {
	api := Router.Group("order")
	{
		api.Use(middleware.CORSMiddleware())

		api.GET(":id", services.GetOrderById())
		api.GET("", services.GetAllOrders())
		api.POST(":id", services.CompleteOrder())
	}
}
