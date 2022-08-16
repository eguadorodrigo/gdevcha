package routes

import (
	"gdevcha/main/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {
	route.POST("/usuarios", controller.CreateUsers)
	route.GET("/usuarios", controller.RetrieveUsers)
	route.PUT("/usuarios/:id", controller.UpdateUsers)
	route.DELETE("/usuarios/:id", controller.DeleteUsers)
}
