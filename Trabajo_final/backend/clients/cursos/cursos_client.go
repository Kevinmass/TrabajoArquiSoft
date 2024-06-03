package cursos

import (
	"backend/domain/users"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCursosTotales() ([]users.CursosData, error) {
	var cursos []users.CursosData
	err := Db.Find(&cursos).Error
	if err != nil {
		log.Error("Error getting cursos: ", err)
		return nil, err
	}
	return cursos, nil
}

func GetCursoPorID(id uint) (*users.CursosData, error) {
	curso := &users.CursosData{}
	err := Db.First(curso, id).Error
	if err != nil {
		log.Error("Error getting curso: ", err)
		return nil, err
	}
	return curso, nil
}

func POSTcrearCurso(curso *users.CursosData) error {

	//Si el curso ya existe, no se puede crear
	var cursoDB users.CursosData
	err1 := Db.Where("nombre = ?", curso.Nombre).First(&cursoDB).Error
	if err1 == nil {
		log.Error("Error creating curso: ", err1)
		return err1
	}

	//Si el id es duplicado se incrementa en 1
	var cursoDB2 users.CursosData
	err2 := Db.Where("id = ?", curso.ID).First(&cursoDB2).Error
	if err2 == nil {
		curso.ID = curso.ID + 1
	}

	err := Db.Create(curso).Error
	if err != nil {
		log.Error("Error creating curso: ", err)
		return err
	}
	return nil
}

func EliminarCurso(curso *users.CursosData) error {

	//con solo el id del curso, se puede eliminar
	err := Db.Delete(curso).Error
	if err != nil {
		log.Error("Error deleting curso: ", err)
		return err
	}

	//eliminar las inscripciones del curso
	var inscripciones []users.InscriptosData
	err = Db.Where("curso_id = ?", curso.ID).Find(&inscripciones).Error
	if err != nil {
		log.Error("Error getting inscripciones: ", err)
		return err
	}
	for _, inscripcion := range inscripciones {
		err = Db.Delete(&inscripcion).Error
		if err != nil {
			log.Error("Error deleting inscripcion: ", err)
			return err
		}
	}

	return nil
}

func PUTmodificarCurso(curso *users.CursosData) error {

	//buscar el curso en la base de datos
	var cursoDB users.CursosData
	err := Db.Where("id = ?", curso.ID).First(&cursoDB).Error
	if err != nil {
		log.Error("Error getting curso: ", err)
		return err
	}

	//verificar si ya existe un curso con el mismo nombre
	var cursoDB2 users.CursosData
	err2 := Db.Where("nombre = ?", curso.Nombre).First(&cursoDB2).Error
	if err2 == nil {
		log.Error("Error modifying curso: ", err2)
		return err2
	}

	//modificar el curso, menos el id y la fecha de creacion
	cursoDB.Nombre = curso.Nombre
	cursoDB.Descripcion = curso.Descripcion
	cursoDB.Estado = curso.Estado
	cursoDB.FechaModificacion = curso.FechaModificacion

	err = Db.Save(&cursoDB).Error
	if err != nil {
		log.Error("Error saving curso: ", err)
		return err
	}

	return nil
}

// Funcion is admin se fija si es admin el usiario del login

func IsAdmin(cliente *users.ClientData) bool {

	// buscar en la base de datos si el cliente es admin

	var clienteDB users.ClientData
	err := Db.Where("user = ?", cliente.User).First(&clienteDB).Error
	if err != nil {
		log.Error("Error getting cliente: ", err)
		return false
	}

	return clienteDB.Admin

}

//Buscar los nombres de todos los cursos que posean alguna similitud con la palabra ingresada

func GetCursosPorNombre(nombreDatabase *users.CursosData) ([]users.CursosData, error) {
	nombre := nombreDatabase.Nombre
	var cursos []users.CursosData
	err := Db.Where("nombre LIKE ?", "%"+nombre+"%").Find(&cursos).Error
	if err != nil {
		log.Error("Error getting cursos: ", err)
		return nil, err
	}
	return cursos, nil
}
