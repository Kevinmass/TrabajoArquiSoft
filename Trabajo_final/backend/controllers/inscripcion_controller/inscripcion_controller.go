package inscripcioncontroller

import (
	"backend/dto"
	service "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func POSTinscripcion(c *gin.Context) {
	var (
		clien_cursoID dto.InscriptosDataDto
	)
	if err := c.ShouldBindJSON(&clien_cursoID); err != nil {
		c.JSON(400, "Error al bindear json")
		return
	}
	err := service.InscripcionService.POSTinscripcion(&clien_cursoID)
	if err != nil {
		c.JSON(500, "Error al inscribirse")
		return
	}
	c.JSON(http.StatusCreated, "Inscripcion realizada")
}

func GetcursosdelCliente(c *gin.Context) {

	log.Debug("Cliente user to load: " + c.Param("user"))

	user := c.Param("user")

	relacion := &dto.InscriptosDataDto{
		User: user,
	}

	cursos, RestErr := service.InscripcionService.GetcursosdelCliente(relacion)
	if RestErr != nil {
		c.JSON(RestErr.StatusCode, RestErr.Message)
		return
	}

	c.JSON(http.StatusOK, cursos)

}
