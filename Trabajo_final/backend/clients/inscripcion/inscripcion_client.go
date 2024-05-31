package inscripcion

import (
	"backend/domain/users"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func relacionarClienteCurso(clienteID, cursoID uint) {
	cliente := users.ClientData{}
	curso := users.CursosData{}

	Db.First(&cliente, clienteID)
	Db.First(&curso, cursoID)

	Db.Model(&cliente).Association("Cursos").Append(&curso)

}

func POSTinscripcion(clienteID, cursoID uint) error {
	relacionarClienteCurso(clienteID, cursoID)
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
