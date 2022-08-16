package controller

import (
	"gdevcha/main/config"
	"gdevcha/main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(context *gin.Context) {
	users := []models.Usuario{}
	config.DB.Find(&users)
	context.IndentedJSON(http.StatusOK, &users)
}
