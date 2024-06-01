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

func GetCursoPorID(id int) (*users.CursosData, error) {
	curso := &users.CursosData{}
	err := Db.First(curso, id).Error
	if err != nil {
		log.Error("Error getting curso: ", err)
		return nil, err
	}
	return curso, nil
}

func POSTcrearCurso(curso *users.CursosData) error {
	err := Db.Create(curso).Error
	if err != nil {
		log.Error("Error creating curso: ", err)
		return err
	}
	return nil
}

func EliminarCurso(curso *users.CursosData) error {
	err := Db.Delete(curso).Error
	if err != nil {
		log.Error("Error deleting curso: ", err)
		return err
	}
	return nil
}

func PUTmodificarCurso(curso *users.CursosData) error {
	err := Db.Save(curso).Error
	if err != nil {
		log.Error("Error saving curso): ", err)
		return err
	}
	return nil
}

// Funcion is admin se fija si es admin el usiario del login

func IsAdmin(cliente *users.ClientData) bool {

	// buscar en la base de datos si el cliente es admin

	var clienteDB users.ClientData
	err := Db.Where("user = ?", cliente.Email).First(&clienteDB).Error
	if err != nil {
		log.Error("Error getting cliente: ", err)
		return false
	}

	return clienteDB.Admin

}
