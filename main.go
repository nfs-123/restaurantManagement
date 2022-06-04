package main

import (
	"fmt"
	"os"
	"restaurantManagement/database"
	"restaurantManagement/middleware"
	"restaurantManagement/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("hii")

	router := gin.New()
	routes.UserRoutes(router)
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)

}
