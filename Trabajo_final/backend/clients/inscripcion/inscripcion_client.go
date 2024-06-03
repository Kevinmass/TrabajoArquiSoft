package inscripcion

import (
	"backend/domain/users"

	"github.com/jinzhu/gorm"

	e "backend/utils"
)

var Db *gorm.DB

func POSTinscripcion(clienteID, cursoID uint) e.RestErr {

	// verificar si existe el cliente y el curso
	cliente := users.ClientData{}
	curso := users.CursosData{}
	Db.First(&cliente, clienteID)
	Db.First(&curso, cursoID)
	if cliente.ID == 0 || curso.ID == 0 {
		return e.RestErr{Message: "Cliente o curso no encontrado", StatusCode: 404}
	}

	// verificar si el cliente ya esta inscrito en el curso
	var inscripcion users.InscriptosData
	Db.Where("cliente_id = ? AND curso_id = ?", clienteID, cursoID).First(&inscripcion)
	if inscripcion.ID != 0 {
		return e.RestErr{Message: "Cliente ya inscrito en el curso", StatusCode: 400}
	}

	// crear la inscripcion
	inscripcion = users.InscriptosData{ClienteID: clienteID, CursoID: cursoID}
	Db.Create(&inscripcion)
	return e.RestErr{Message: "Inscripcion exitosa", StatusCode: 200}

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
