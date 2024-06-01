package inscripcion

import (
	"backend/domain/users"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func POSTinscripcion(clienteID, cursoID uint) error {

	realacion := users.InscriptosData{}
	realacion.ClienteID = clienteID
	realacion.CursoID = cursoID
	Db.Create(&realacion)

	return nil

}

func GetcursosdelCliente(clienteID uint) ([]users.CursosData, error) {
	var cursos []users.CursosData
	cliente := users.ClientData{}
	Db.First(&cliente, clienteID)
	err := Db.Model(&cliente).Related(&cursos, "Cursos").Error
	if err != nil {
		return nil, err
	}
	return cursos, nil
}
