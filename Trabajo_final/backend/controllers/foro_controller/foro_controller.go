package forocontroller

import (
	"backend/dto"
	service "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func POSTcomentario(c *gin.Context) {

	var comentario *dto.ForodataDto
	if err := c.ShouldBindJSON(&comentario); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.ForoService.POSTcomentario(comentario)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, "Comentario creado")

}

func GETcomentarios(c *gin.Context) {

	var forodata *dto.ForodataDto
	if err := c.ShouldBindJSON(&forodata); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comentarios, err := service.ForoService.GETcomentarios(forodata)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, comentarios)

}

func DELETEcomentario(c *gin.Context) {

	var forodata *dto.ForodataDto
	if err := c.ShouldBindJSON(&forodata); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.ForoService.DELETEcomentario(forodata)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, "Comentario eliminado")

}
