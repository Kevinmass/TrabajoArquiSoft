package controllers

import (
	"log"
	"net/http"

	services "backend/services/users"

	"github.com/gin-gonic/gin"
)

func POSTArchivos(c *gin.Context) {

	file, err := c.FormFile("file")

	log.Println(file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	cursoId := c.Param("id")
	log.Println(cursoId)
	if restErr := services.FileUploadService.SubirArchivos(file, cursoId); restErr != nil {
		c.JSON(restErr.StatusCode, gin.H{"error": restErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func GetArchivos(c *gin.Context) {
	cursoId := c.Param("id")
	log.Println(cursoId)
	archivos, restErr := services.FileService.GetArchivos(cursoId)
	if restErr != nil {
		c.JSON(restErr.StatusCode, gin.H{"error": restErr.Message})
		return
	}
	c.JSON(http.StatusOK, archivos)
}
