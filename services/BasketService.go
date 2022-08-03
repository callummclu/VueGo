package services

import (
	"go-vue-ecommerce-site/configs"
	"go-vue-ecommerce-site/helpers"
	"go-vue-ecommerce-site/models"
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
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func GetBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var basket models.Basket
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		err = basketCollection.FindOne(c, bson.D{{"_id", objectId}}).Decode(&basket)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
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
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		}

		filter := bson.D{{"_id", objectId}}
		update := bson.D{{"items", basket.Items}}

		result, err := basketCollection.ReplaceOne(c, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func AddToBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var currentBasket models.Basket
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "cannot convert id to int"})
			return
		}

		err = basketCollection.FindOne(c, bson.D{{"_id", objectId}}).Decode(&currentBasket)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "cannot find item"})
			return
		}

		var product models.Product
		c.BindJSON(&product)

		if len(currentBasket.Items) > 0 {
			for i, item := range currentBasket.Items {
				if item.ID == product.ID {
					currentBasket.Items[i].Quantity++
					currentBasket.Items[i].Price += product.Price
					break
				} else {
					if item == currentBasket.Items[len(currentBasket.Items)-1] {
						currentBasket.Items = append(currentBasket.Items, product)
					}
				}
			}
		} else {
			currentBasket.Items = append(currentBasket.Items, product)
		}

		basketTotal := 0.00

		for _, item := range currentBasket.Items {
			basketTotal += float64(item.Price)
		}

		filter := bson.D{{"_id", objectId}}
		update := bson.D{{"items", currentBasket.Items}, {"total", basketTotal}}

		result, err := basketCollection.ReplaceOne(c, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot replace item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func DeleteOneFromBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var currentBasket models.Basket
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "cannot convert id to int"})
			return
		}

		err = basketCollection.FindOne(c, bson.D{{"_id", objectId}}).Decode(&currentBasket)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "cannot find item"})
			return
		}

		var product models.Product
		c.BindJSON(&product)

		if len(currentBasket.Items) > 0 {
			if len(currentBasket.Items) == 1 {
				currentBasket.Items = nil
			}

			for i, item := range currentBasket.Items {
				if item.ID == product.ID {
					helpers.RemoveIndex(currentBasket.Items, i)
				}
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "there are not items in basket."})
		}

		basketTotal := 0.00

		for _, item := range currentBasket.Items {
			basketTotal += float64(item.Price)
		}

		filter := bson.D{{"_id", objectId}}
		update := bson.D{{"items", currentBasket.Items}, {"total", basketTotal}}

		result, err := basketCollection.ReplaceOne(c, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot replace item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
