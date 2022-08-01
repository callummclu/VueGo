package main

import (
	"go-vue-ecommerce-site/controllers"
	"go-vue-ecommerce-site/models"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	models.Connect()
	controllers.BaseController()
	controllers.Router.Run(":8000")
}
