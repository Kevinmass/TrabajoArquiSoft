package services

import (
	inscripcionClient "backend/clients/inscripcion"
	"backend/dto"
)

type inscripcionService struct{}

type inscripcionServiceInterface interface {
	POSTinscripcion(relacionU *dto.InscriptosDataDto) error
	GetcursosdelCliente(relacion *dto.InscriptosDataDto) ([]dto.CursosDatadto, error)
}

var (
	InscripcionService inscripcionServiceInterface
)

func init() {
	InscripcionService = &inscripcionService{}
}

func (s *inscripcionService) POSTinscripcion(relacionU *dto.InscriptosDataDto) error {

	err := inscripcionClient.POSTinscripcion(relacionU.ClienteID, relacionU.CursoID)
	if err != nil {
		return err
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
			ID:            curso.ID,
			Nombre:        curso.Nombre,
			Descripcion:   curso.Descripcion,
			FechaCreacion: curso.FechaCreacion,
			Estado:        curso.Estado,
		}
	}

	return cursosDto, nil

}
