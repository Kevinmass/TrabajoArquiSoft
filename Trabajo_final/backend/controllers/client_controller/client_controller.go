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

func POSTlogin(c *gin.Context) {

	var cliente *dto.ClientDatadto
	if err := c.ShouldBindJSON(&cliente); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.ClientService.POSTlogin(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	//guardar en cookie
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	//guardar usuario en cookie
	c.SetCookie("user", cliente.User, 3600, "/", "localhost", false, true)

}
