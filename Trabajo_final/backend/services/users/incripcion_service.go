package services

import (
	inscripcionClient "backend/clients/inscripcion"
	"backend/dto"
	"backend/utils"
)

type inscripcionService struct{}

type inscripcionServiceInterface interface {
	POSTinscripcion(relacionU *dto.InscriptosDataDto) *utils.RestErr
	GetcursosdelCliente(relacion *dto.InscriptosDataDto) ([]dto.CursosDatadto, error)
}

var (
	InscripcionService inscripcionServiceInterface
)

func init() {
	InscripcionService = &inscripcionService{}
}

func (s *inscripcionService) POSTinscripcion(relacionU *dto.InscriptosDataDto) *utils.RestErr {

	err := inscripcionClient.POSTinscripcion(relacionU.ClienteID, relacionU.CursoID)

	if err.StatusCode != 200 {
		return &utils.RestErr{Message: "Error al inscribirse", StatusCode: err.StatusCode}
	}

	return nil
}

func (s *inscripcionService) GetcursosdelCliente(relacion *dto.InscriptosDataDto) ([]dto.CursosDatadto, error) {

	cursos, err := inscripcionClient.GetcursosdelCliente(relacion.ClienteID)
	if err != nil {
		return nil, err
	}

	var cursosDto []dto.CursosDatadto = make([]dto.CursosDatadto, len(cursos))

	for i, curso := range cursos {
		cursosDto[i] = dto.CursosDatadto{
			ID:          curso.ID,
			Nombre:      curso.Nombre,
			Descripcion: curso.Descripcion,
			Estado:      curso.Estado,
		}
	}

	return cursosDto, nil

}
