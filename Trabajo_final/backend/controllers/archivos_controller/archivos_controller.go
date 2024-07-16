package controllers

import (
	"log"
	"net/http"

	services "backend/services/users"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	log.Println(file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	if restErr := services.FileUploadService.UploadFile(file); restErr != nil {
		c.JSON(restErr.StatusCode, gin.H{"error": restErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
