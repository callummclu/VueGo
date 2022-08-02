package services

import (
	"go-vue-ecommerce-site/configs"
	"go-vue-ecommerce-site/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var basketCollection *mongo.Collection = configs.GetCollection(configs.DB, "baskets")

func CreateBasket() gin.HandlerFunc {

	return func(c *gin.Context) {
		var basket models.Basket
		c.BindJSON(&basket)
		result, err := basketCollection.InsertOne(c, basket)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func GetBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var basket models.Basket
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		err = basketCollection.FindOne(c, bson.D{{"_id", objectId}}).Decode(&basket)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": basket})
	}
}

func UpdateBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var basket models.Basket
		c.BindJSON(&basket)

		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		filter := bson.D{{"_id", objectId}}
		update := bson.D{{"items", basket.Items}}

		result, err := basketCollection.ReplaceOne(c, filter, update)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
