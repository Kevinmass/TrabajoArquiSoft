package inscripcioncontroller

import (
	"backend/dto"
	service "backend/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	var (
		clienteID *dto.InscriptosDataDto
	)
	clienteIDInt, _ := strconv.Atoi(c.Param("clienteID"))
	clienteID = &dto.InscriptosDataDto{ID: clienteIDInt}

	cursos, err := service.InscripcionService.GetcursosdelCliente(clienteID)
	if err != nil {
		c.JSON(500, "Error al obtener los cursos")
		return
	}
	c.JSON(http.StatusOK, cursos)
}
