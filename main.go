package main

import (
	"gdevcha/main/config"
	"gdevcha/main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":9090")
}
