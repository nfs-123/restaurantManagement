package routes

import (
	"github.com/gin-gonic/gin"
	"restaurantManagement/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.Getmenus())
	incomingRoutes.GET("/menus/:menu_id", controllers.Getmenu())
	incomingRoutes.POST("/menus", controllers.Createmenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.Updatemenu())

}
