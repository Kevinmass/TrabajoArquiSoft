package cursos

import (
	"backend/clients/client"
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

	clienteLOG := client.GETloginCliente(cliente)

	err := Db.Find(&clienteLOG, "admin = ?", true).Error
	if err == nil {

		return true
	}
	return false

}
