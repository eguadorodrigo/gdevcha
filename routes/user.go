package routes

import (
	"gdevcha/main/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {
	route.GET("/usuarios", controller.UserController)
}
