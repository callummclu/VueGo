package main

import (
	"go-vue-ecommerce-site/configs"
	"go-vue-ecommerce-site/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	configs.ConnectDB()
	controllers.BaseController()
}
