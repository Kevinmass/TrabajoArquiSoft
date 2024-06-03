package app

import (
	clientController "backend/controllers/client_controller"
	cursosController "backend/controllers/cursos_controller"
	inscripcionController "backend/controllers/inscripcion_controller"
)

func mapUrls() {

	router.GET("/login", clientController.GETloginCliente)
	router.POST("/register", clientController.POSTregistro)
	router.GET("/isAdmin", cursosController.IsAdmin)
	router.GET("/cursos", cursosController.GetCursosTotales)
	router.GET("/cursos/:id", cursosController.GetCursoPorID)
	router.POST("/crearCurso", cursosController.POSTcrearCurso)
	router.POST("/inscribirse", inscripcionController.POSTinscripcion)
	router.POST("/eliminarCurso", cursosController.EliminarCurso)
	router.PUT("/modificarCurso", cursosController.PUTmodificarCurso)
	router.GET("/cursosPorNombre", cursosController.GetCursosPorNombre)
}
