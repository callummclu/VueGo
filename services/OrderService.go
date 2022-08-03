package services

import (
	"go-vue-ecommerce-site/configs"
	"go-vue-ecommerce-site/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = configs.GetCollection(configs.DB, "orders")

func GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		err = orderCollection.FindOne(c, bson.D{{"_id", objectId}}).Decode(&order)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": order})
	}
}

func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		var orders []models.Order

		cursor, err := orderCollection.Find(c, bson.D{})
		if err != nil {
			defer cursor.Close(c)
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		}

		for cursor.Next(c) {
			err := cursor.Decode(&order)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err})
			}
			orders = append(orders, order)
		}
		c.JSON(http.StatusOK, gin.H{"data": orders})
	}
}

func CompleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		filter := bson.D{{"_id", objectId}}
		update := bson.D{{"status", "SHIPPED"}}

		result, err := orderCollection.ReplaceOne(c, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot change order status."})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})

	}
}
