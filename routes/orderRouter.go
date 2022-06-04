package routes

import (
	"restaurantManagement/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.Getorders())
	incomingRoutes.GET("/orders/:order_id", controllers.Getorder())
	incomingRoutes.POST("/orders", controllers.Createorder())
	incomingRoutes.PATCH("/orders/:order_id", controllers.Updateorder())

}
