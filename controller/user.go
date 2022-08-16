package controller

import (
	"gdevcha/main/config"
	"gdevcha/main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RetrieveUsers(context *gin.Context) {
	users := []models.Usuario{}
	config.DB.Find(&users)
	context.IndentedJSON(http.StatusOK, &users)
}


func CreateUsers(context *gin.Context) {
	var usuario models.Usuario
	context.BindJSON(&usuario)
	config.DB.Create(&usuario)
	context.IndentedJSON(http.StatusCreated, &usuario)
}


func UpdateUsers(context *gin.Context) {
	var usuario models.Usuario
	config.DB.Where("id=?", context.Param("id")).First(&usuario)
	context.BindJSON(&usuario)
	config.DB.Save(&usuario)
	context.IndentedJSON(http.StatusAccepted, &usuario)
}

func DeleteUsers(context *gin.Context) {
	var usuario models.Usuario
	config.DB.Where("id=?", context.Param("id")).Delete(&usuario)
	context.IndentedJSON(http.StatusGone, &usuario)
}