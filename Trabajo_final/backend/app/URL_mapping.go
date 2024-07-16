package app

import (
	archivoController "backend/controllers/archivos_controller"
	clientController "backend/controllers/client_controller"
	cursosController "backend/controllers/cursos_controller"
	foroController "backend/controllers/foro_controller"
	inscripcionController "backend/controllers/inscripcion_controller"
)

func mapUrls() {

	router.POST("/register", clientController.POSTregistro)
	router.POST("/login", clientController.POSTlogin)
	router.GET("/isAdmin/:user", cursosController.IsAdmin)
	router.POST("/crearCurso", cursosController.POSTcrearCurso)
	router.POST("/inscribirse", inscripcionController.POSTinscripcion)
	router.GET("/cursos", cursosController.GetCursosTotales)
	router.GET("/cursos/:id", cursosController.GetCursoPorID)
	router.POST("/eliminarCurso", cursosController.EliminarCurso)
	router.PUT("/modificarCurso", cursosController.PUTmodificarCurso)
	router.GET("/cursosPorNombre/:nombre", cursosController.GetCursosPorNombre)
	router.GET("/GetCursosdelCliente/:user", inscripcionController.GetcursosdelCliente)
	router.POST("/POSTcomentario", foroController.POSTcomentario)
	router.GET("/GETcomentarios", foroController.GETcomentarios)
	router.DELETE("/DELETEcomentario", foroController.DELETEcomentario)
	router.POST("/POSTarchivo/:id", archivoController.POSTArchivos)
	router.GET("/GETarchivo/:id", archivoController.GetArchivos)
}
