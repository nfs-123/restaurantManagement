package controllers

import (
	"context"
	"fmt"
	"net/http"
	"restaurantManagement/database"
	"restaurantManagement/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"foodId": foodId}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while fetching food item..",
			})
		}
		c.JSON(http.StatusOK, food)
		defer cancel()

	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := c.Bind(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		validationError := validate.Struct(food)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError.Error(),
			})
			return

		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": msg,
			})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(food.Price, 2)
		food.Price = num
		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Food item not created..",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

// func Sprintf(s string) {
// 	panic("unimplemented")
// }
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func round(num float64) int {

}
func toFixed(num float64, precision int) float64 {

}
