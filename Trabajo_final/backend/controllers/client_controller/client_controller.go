package clientcontroller

import (
	"backend/dto"
	service "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func POSTregistro(c *gin.Context) {
	var cliente *dto.ClientDatadto
	if err := c.ShouldBindJSON(&cliente); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cliente, err := service.ClientService.POSTregistro(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, cliente)
}

func PUThashearPassword(c *gin.Context) {
	var cliente *dto.ClientDatadto
	if err := c.ShouldBindJSON(&cliente); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cliente, err := service.ClientService.PUThashearPassword(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, cliente)
}

func GETverificarCliente(c *gin.Context) {
	var cliente *dto.ClientDatadto
	if err := c.ShouldBindJSON(&cliente); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cliente, err := service.ClientService.GETverificarCliente(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, cliente)
}

func GETloginCliente(c *gin.Context) {
	var cliente *dto.ClientDatadto
	if err := c.ShouldBindJSON(&cliente); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cliente, err := service.ClientService.GETloginCliente(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, cliente)
	log.Info("Logiado como cliente: ", cliente.User)
}
