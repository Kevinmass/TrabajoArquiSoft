package services

import (
	inscripcionClient "backend/clients/inscripcion"
	"backend/dto"
	e "backend/utils"
)

type inscripcionService struct{}

type inscripcionServiceInterface interface {
	POSTinscripcion(relacionU *dto.InscriptosDataDto) *e.RestErr
	GetcursosdelCliente(relacion *dto.InscriptosDataDto) ([]dto.CursosDatadto, *e.RestErr)
	GetUserID(user string) uint
}

var (
	InscripcionService inscripcionServiceInterface
)

func init() {
	InscripcionService = &inscripcionService{}
}

func (s *inscripcionService) GetUserID(user string) uint {

	ClienteID := inscripcionClient.GetUserID(user)

	return ClienteID
}

func (s *inscripcionService) POSTinscripcion(relacionU *dto.InscriptosDataDto) *e.RestErr {

	ClienteID := s.GetUserID(relacionU.User)

	err := inscripcionClient.POSTinscripcion(ClienteID, relacionU.CursoID)

	if err.StatusCode != 200 {
		return &e.RestErr{Message: "Error al inscribirse", StatusCode: err.StatusCode}
	}

	return nil
}

func (s *inscripcionService) GetcursosdelCliente(relacion *dto.InscriptosDataDto) ([]dto.CursosDatadto, *e.RestErr) {

	ClienteID := s.GetUserID(relacion.User)

	cursos, err := inscripcionClient.GetcursosdelCliente(ClienteID)

	if err.StatusCode != 200 {
		return nil, &e.RestErr{Message: "Error al obtener los cursos", StatusCode: err.StatusCode}
	}

	var cursosDto []dto.CursosDatadto
	var err2 *e.RestErr
	err2 = nil

	for _, cursos := range cursos {
		cursoDto := &dto.CursosDatadto{
			ID:          cursos.ID,
			Nombre:      cursos.Nombre,
			Descripcion: cursos.Descripcion,
			User:        cursos.User,
			Duracion:    cursos.Duracion,
			Requisitos:  cursos.Requisitos,
		}
		cursosDto = append(cursosDto, *cursoDto)
	}

	if len(cursosDto) == 0 {
		err2 = &e.RestErr{StatusCode: 404}
		cursosDto = []dto.CursosDatadto{}
	}

	return cursosDto, err2
}
