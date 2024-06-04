package clientcontroller

import (
	"backend/dto"
	service "backend/services/users"
	"net/http"

	"github.com/gin-contrib/sessions"
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
	isValid := service.ClientService.POSTlogin(cliente)
	if isValid == false {
		c.JSON(http.StatusUnauthorized, "Usuario o contraseña incorrectos")
		return
	}

	c.JSON(http.StatusOK, "Usuario logueado")
}

func GETvalidar(c *gin.Context) {
	var cliente *dto.ClientDatadto

	if err := c.ShouldBindJSON(&cliente); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cliente, err := service.ClientService.GETvalidar(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, cliente)

	// guardar nombre de usuario en la sesion
	session := sessions.Default(c)
	session.Set("user", cliente.User)
	session.Save()

}
