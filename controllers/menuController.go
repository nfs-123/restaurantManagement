package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurantManagement/models"
	"time"

	// "restaurantManagement/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func Getmenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		result, err := menuCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while getting menus",
			})
			return
		}
		defer cancel()
		var allMenus []bson.M
		if err := result.All(context.TODO(), &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)

	}
}
func Getmenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menuId": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while fetching menu ..",
			})
		}
		c.JSON(http.StatusOK, menu)
		defer cancel()

	}
}
func Createmenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		if err := c.Bind(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		validationError := validate.Struct(menu)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError.Error(),
			})
			return

		}
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()
		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Menu item not created..",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}
func Updatemenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
