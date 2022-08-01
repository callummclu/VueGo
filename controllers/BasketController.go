package controllers

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func BasketController() {
	api := Router.Group("basket")
	{
		api.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"basket": "",
			})
		})
	}

	/*
		CREATE BASKET FOR SESSION
		ALLOW USER TO ADD ITEMS TO BASKET
	*/
}
