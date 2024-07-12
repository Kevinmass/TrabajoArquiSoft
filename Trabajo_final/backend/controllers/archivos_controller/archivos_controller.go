package controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateAndWriteFile(c *gin.Context) {
	log.Println("Creating a file...")

	// Define the complete file path
	dirPath := "Trabajo_final/backend/archivos"
	filePath := dirPath + "/mi_archivo.txt"
	log.Println(filePath)

	// Ensure the directory exists
	err := os.MkdirAll(dirPath, os.ModePerm)
	log.Println("err0")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create directory: " + err.Error(),
		})
		return
	}
	log.Println("err1")
	// Create the file
	file, err := os.Create(filePath)
	log.Println("err2")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create file: " + err.Error(),
		})
		return
	}
	log.Println("err03")
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Failed to close the file:", err)
		}
	}()
	log.Println("File created")

	// Write the content to the file
	content, err := c.GetRawData() // Assuming 'content' is in the request body
	log.Println(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read request body: " + err.Error(),
		})
		return
	}
	log.Println("err4")
	_, err = file.Write(content)
	log.Println("err5")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to write to file: " + err.Error(),
		})
		return
	}
	log.Println("err6")
	c.JSON(http.StatusOK, gin.H{
		"message": "File created and written successfully",
	})
}

func ReadFile(c *gin.Context) {
	path := c.Param("path") // Assuming 'path' is a URL parameter

	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", bytes)
}
