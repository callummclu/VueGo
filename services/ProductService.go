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

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		var products []models.Product

		cursor, err := productCollection.Find(c, bson.D{})
		if err != nil {
			defer cursor.Close(c)
			log.Fatal(err)
		}

		for cursor.Next(c) {
			err := cursor.Decode(&product)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, product)
		}
		c.JSON(http.StatusOK, gin.H{"data": products})
	}
}

func GetOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		err = productCollection.FindOne(c, bson.D{{"_id", objectId}}).Decode(&product)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": product})
	}
}

func CreateOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		c.BindJSON(&product)
		result, err := productCollection.InsertOne(c, product)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func DeleteOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		_, err = productCollection.DeleteOne(c, bson.D{{"_id", objectId}})
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})

	}
}

func EditOneProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		c.BindJSON(&product)

		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		filter := bson.D{{"_id", objectId}}
		update := bson.D{{"productName", product.ProductName}, {"price", product.Price}, {"quantity", product.Quantity}}

		result, err := productCollection.ReplaceOne(c, filter, update)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
