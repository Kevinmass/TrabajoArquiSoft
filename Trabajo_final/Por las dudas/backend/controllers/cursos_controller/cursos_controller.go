package cursoscontroller

import (
	"backend/dto"
	service "backend/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	e "backend/utils"
)

func IsAdmin(c *gin.Context) {
	var cliente *dto.CursoClienteDataDto
	if err := c.ShouldBindJSON(&cliente); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, bool := service.CursoService.IsAdmin(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, bool)
}

func GetCursosTotales(c *gin.Context) {
	cursos, err := service.CursoService.GetCursosTotales()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, cursos)
}

func GetCursoPorID(c *gin.Context) {

	log.Debug("Curso id to load: " + c.Param("id"))

	// convertir el id a uint
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err1 *e.RestErr

	curso, err1 := service.CursoService.GetCursoPorID(uint(id))
	if err1.StatusCode != 0 {
		c.JSON(err1.StatusCode, err1)
		return
	}

	c.JSON(http.StatusOK, curso)

}

func POSTcrearCurso(c *gin.Context) {

	var (
		cursoV *dto.CursoClienteDataDto
	)
	// requiero un multipart form
	if err := c.ShouldBindJSON(&cursoV); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	curso, err := service.CursoService.POSTcrearCurso(cursoV)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, curso)

}

func PUTmodificarCurso(c *gin.Context) {

	var (
		cursoV *dto.CursoClienteDataDto
	)

	if err := c.ShouldBindJSON(&cursoV); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	curso, err := service.CursoService.PUTmodificarCurso(cursoV)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, curso)
}

func EliminarCurso(c *gin.Context) {

	var (
		cursoV *dto.CursoClienteDataDto
	)

	if err := c.ShouldBindJSON(&cursoV); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CursoService.EliminarCurso(cursoV)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, "Curso eliminado")

}

func GetCursosPorNombre(c *gin.Context) {

	var (
		nombre *dto.CursoBusqueda
	)
	if err := c.ShouldBindJSON(&nombre); err != nil {
		log.Error("Error binding json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cursos, err := service.CursoService.GetCursosPorNombre(nombre)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, cursos)
}
