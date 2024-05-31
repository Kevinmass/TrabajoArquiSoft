package services

import (
	cursoClient "backend/clients/cursos"
	"backend/domain/users"
	"backend/dto"
	e "backend/utils"
)

type cursoService struct{}

type cursoServiceInterface interface {
	POSTcrearCurso(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr)
	IsAdmin(cliente *dto.ClientAdminto) *e.RestErr
	GetCursoPorID(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr)
	GetCursosTotales() ([]dto.CursosDatadto, *e.RestErr)
	PUTmodificarCurso(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr)
	EliminarCurso(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr)
}

var (
	CursoService cursoServiceInterface
)

func init() {
	CursoService = &cursoService{}
}

// si el cliente es admin crear curso

func (s *cursoService) IsAdmin(cliente *dto.ClientAdminto) *e.RestErr {
	if !cliente.Admin {
		return &e.RestErr{Message: "No es admin", StatusCode: 400}
	}
	return nil
}

func (s *cursoService) POSTcrearCurso(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr) {

	cursoData := &users.CursosData{
		Nombre:        curso.Nombre,
		Descripcion:   curso.Descripcion,
		FechaCreacion: curso.FechaCreacion,
		Estado:        curso.Estado,
	}

	CheckAdmin := s.IsAdmin(&curso.UsuarioAdmin)

	if CheckAdmin != nil {
		return nil, CheckAdmin
	}

	err := cursoClient.POSTcrearCurso(cursoData)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return curso, nil

}

func (s *cursoService) GetCursoPorID(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr) {

	cursoData := &users.CursosData{
		ID: curso.ID,
	}

	cursoData, err := cursoClient.GetCursoPorID(cursoData.ID)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return curso, nil
}

func (s *cursoService) GetCursosTotales() ([]dto.CursosDatadto, *e.RestErr) {
	cursos, err := cursoClient.GetCursosTotales()
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	var cursosDto []dto.CursosDatadto
	for _, curso := range cursos {
		cursoDto := &dto.CursosDatadto{
			ID:                curso.ID,
			Nombre:            curso.Nombre,
			Descripcion:       curso.Descripcion,
			FechaCreacion:     curso.FechaCreacion,
			FechaModificacion: curso.FechaModificacion,
			Estado:            curso.Estado,
		}
		cursosDto = append(cursosDto, *cursoDto)
	}

	return cursosDto, nil
}

func (s *cursoService) PUTmodificarCurso(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr) {
	cursoData := &users.CursosData{
		ID:                curso.ID,
		Nombre:            curso.Nombre,
		Descripcion:       curso.Descripcion,
		FechaModificacion: curso.FechaModificacion,
		Estado:            curso.Estado,
	}

	CheckAdmin := s.IsAdmin(&curso.UsuarioAdmin)

	if CheckAdmin != nil {
		return nil, CheckAdmin
	}

	err := cursoClient.PUTmodificarCurso(cursoData)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return curso, nil
}

func (s *cursoService) EliminarCurso(curso *dto.CursosDatadto) (*dto.CursosDatadto, *e.RestErr) {
	cursoData := &users.CursosData{
		ID: curso.ID,
	}

	CheckAdmin := s.IsAdmin(&curso.UsuarioAdmin)

	if CheckAdmin != nil {
		return nil, CheckAdmin
	}

	err := cursoClient.PUTmodificarCurso(cursoData)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return curso, nil
}