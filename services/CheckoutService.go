package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckoutBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		result, err := basketCollection.DeleteOne(c, bson.D{{"_id", objectId}})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		// ADD SOME MORE CHECKOUT THINGS IN HERE

		c.JSON(http.StatusOK, gin.H{"message": "Checkout successful!", "data": result})
	}
}
