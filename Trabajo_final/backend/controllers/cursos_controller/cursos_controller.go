package cursoscontroller

import (
	"backend/dto"
	service "backend/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func IsAdmin(c *gin.Context) {
	user := c.Param("user")
	if user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User parameter is required"})
		return
	}

	cliente := &dto.CursoClienteDataDto{User: user}
	err, isAdmin := service.CursoService.IsAdmin(cliente)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, isAdmin)
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

	curso, RestErr := service.CursoService.GetCursoPorID(uint(id))
	if RestErr != nil {
		c.JSON(RestErr.StatusCode, RestErr.Message)
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

	// Combine the user and curso ID
	userCurso := &dto.InscriptosDataDto{
		User:    cursoV.User,
		CursoID: curso.ID,
	}
	log.Debug(userCurso.User)
	log.Debug(userCurso.CursoID)

	// Call the service function POSTinscripcion to inscribe the user that creates the curso
	err = service.InscripcionService.POSTinscripcion(userCurso)
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

	nombre := c.Param("nombre")

	cursos, RestErr := service.CursoService.GetCursosPorNombre(nombre)

	if RestErr != nil {
		c.JSON(RestErr.StatusCode, RestErr.Message)
		return
	}
	c.JSON(http.StatusOK, cursos)
}

func GetTotalCursos(c *gin.Context) {
	totalCursos, err := service.CursoService.GetTotalCursos()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, totalCursos)
}
