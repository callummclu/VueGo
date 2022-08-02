package controllers

import (
	"go-vue-ecommerce-site/middleware"
	"go-vue-ecommerce-site/services"
)

func ProductContoller() {
	api := Router.Group("product")
	{

		api.Use(middleware.CORSMiddleware())

		api.GET("", services.GetAllProducts())
		api.GET(":id", services.GetOneProduct())
		api.POST("", services.CreateOneProduct())
		api.DELETE("", services.DeleteOneProduct())
		api.PATCH(":id", services.EditOneProduct())
	}
}
