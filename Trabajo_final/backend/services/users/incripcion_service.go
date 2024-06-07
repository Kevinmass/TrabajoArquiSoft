package services

import (
	inscripcionClient "backend/clients/inscripcion"
	"backend/dto"
	"backend/utils"
)

type inscripcionService struct{}

type inscripcionServiceInterface interface {
	POSTinscripcion(relacionU *dto.InscriptosDataDto) *utils.RestErr
	GetcursosdelCliente(relacion *dto.InscriptosDataDto) ([]dto.CursosDatadto, *utils.RestErr)
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

func (s *inscripcionService) POSTinscripcion(relacionU *dto.InscriptosDataDto) *utils.RestErr {

	ClienteID := s.GetUserID(relacionU.User)

	err := inscripcionClient.POSTinscripcion(ClienteID, relacionU.CursoID)

	if err.StatusCode != 200 {
		return &utils.RestErr{Message: "Error al inscribirse", StatusCode: err.StatusCode}
	}

	return nil
}

func (s *inscripcionService) GetcursosdelCliente(relacion *dto.InscriptosDataDto) ([]dto.CursosDatadto, *utils.RestErr) {

	ClienteID := s.GetUserID(relacion.User)

	cursos, err := inscripcionClient.GetcursosdelCliente(ClienteID)

	if err.StatusCode != 200 {
		return nil, &utils.RestErr{Message: "Error al obtener los cursos", StatusCode: err.StatusCode}
	}

	// Convert the type of cursos from []users.CursosData to []dto.CursosDatadto
	var cursosDto []dto.CursosDatadto

	for _, cursos := range cursos {
		cursoDto := &dto.CursosDatadto{
			ID:                cursos.ID,
			Nombre:            cursos.Nombre,
			Descripcion:       cursos.Descripcion,
			ProfesorNombre:    cursos.ProfesorNombre,
			ProfesorApellido:  cursos.ProfesorApellido,
			ProfesorCorreo:    cursos.ProfesorCorreo,
			FechaCreacion:     cursos.FechaCreacion,
			FechaModificacion: cursos.FechaModificacion,
		}
		cursosDto = append(cursosDto, *cursoDto)
	}

	return cursosDto, nil
}
