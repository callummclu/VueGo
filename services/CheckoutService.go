package services

import (
	"go-vue-ecommerce-site/models"
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

		var basket models.Basket
		err = basketCollection.FindOne(c, bson.D{{"_id", objectId}}).Decode(&basket)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "basket not found"})
		}

		order := models.Order{
			BasketContent: basket,
			Status:        "PAID",
		}

		result, err := orderCollection.InsertOne(c, order)

		deleteResult, err := basketCollection.DeleteOne(c, bson.D{{"_id", objectId}})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Checkout successful!", "data": result, "deleteData": deleteResult})
	}
}
