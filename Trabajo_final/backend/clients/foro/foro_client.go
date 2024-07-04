package foro

import (
	"backend/domain/users"

	"github.com/jinzhu/gorm"

	e "backend/utils"
)

var Db *gorm.DB

func POSTcomentario(cursoID uint, user, comment string) e.RestErr {

	// verificar si existe el curso
	curso := users.CursosData{}
	Db.First(&curso, cursoID)
	if curso.ID == 0 {
		return e.RestErr{Message: "Curso no encontrado", StatusCode: 404}
	}

	// crear el comentario
	comentario := users.Forodata{CursoID: cursoID, User: user, Comentario: comment}
	Db.Create(&comentario)
	return e.RestErr{Message: "Comentario creado", StatusCode: 200}

}

func GETcomentarios(cursoID uint) ([]users.Forodata, e.RestErr) {

	// verificar si existe el curso
	curso := users.CursosData{}
	Db.First(&curso, cursoID)
	if curso.ID == 0 {
		return nil, e.RestErr{Message: "Curso no encontrado", StatusCode: 404}
	}

	// obtener los comentarios del curso
	var comentarios []users.Forodata
	Db.Where("curso_id = ?", cursoID).Find(&comentarios)
	return comentarios, e.RestErr{Message: "Comentarios obtenidos", StatusCode: 200}

}

// eliminar comentario si el usuario es el autor

func DELETEcomentario(user string) e.RestErr {

	// verificar si el comentario es del usuario
	comentario := users.Forodata{}
	Db.Where("user = ?", user).First(&comentario)

	if comentario.User != user {
		return e.RestErr{Message: "No hay comentario a eliminar", StatusCode: 404}
	}

	// eliminar el comentario
	Db.Delete(&comentario)
	return e.RestErr{Message: "Comentario eliminado", StatusCode: 200}

}
