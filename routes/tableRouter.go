package routes

import (
	"github.com/gin-gonic/gin"
	"restaurantManagement/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.Gettables())
	incomingRoutes.GET("/tables/:table_id", controllers.Gettable())
	incomingRoutes.POST("/tables", controllers.Createtable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.Updatetable())

}
