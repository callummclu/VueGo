package controllers

import (
	"go-vue-ecommerce-site/configs"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func BaseController() {
	BasketController()
	CheckoutController()
	ProductContoller()
	Router.Run(configs.EnvPORT())
}
