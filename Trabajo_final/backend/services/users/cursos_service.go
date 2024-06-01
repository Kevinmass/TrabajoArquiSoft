package services

import (
	cursoClient "backend/clients/cursos"
	"backend/domain/users"
	"backend/dto"
	e "backend/utils"
)

type cursoService struct{}

type cursoServiceInterface interface {
	POSTcrearCurso(cursoV *dto.CursoClienteDataDto) (*dto.CursoClienteDataDto, *e.RestErr)
	IsAdmin(cliente *dto.ClientDatadto) (*e.RestErr, bool)
	GetCursoPorID(id int) (*dto.CursosDatadto, *e.RestErr)
	GetCursosTotales() ([]dto.CursosDatadto, *e.RestErr)
	PUTmodificarCurso(cursoV *dto.CursoClienteDataDto) (*dto.CursoClienteDataDto, *e.RestErr)
	EliminarCurso(cursoV *dto.CursoClienteDataDto) *e.RestErr
}

var (
	CursoService cursoServiceInterface
)

func init() {
	CursoService = &cursoService{}
}

// si el cliente es admin crear curso

func (s *cursoService) IsAdmin(cliente *dto.ClientDatadto) (*e.RestErr, bool) {

	clienteData := &users.ClientData{
		User: cliente.User,
	}

	CheckAdmin := cursoClient.IsAdmin(clienteData)
	if CheckAdmin == false {
		return &e.RestErr{Message: "No es admin", StatusCode: 400}, false
	}

	CheckAdmin = true

	return nil, CheckAdmin
}

func (s *cursoService) POSTcrearCurso(cursoV *dto.CursoClienteDataDto) (*dto.CursoClienteDataDto, *e.RestErr) {

	cursoData := &users.CursosData{
		ID:            cursoV.ID,
		Nombre:        cursoV.Nombre,
		Descripcion:   cursoV.Descripcion,
		FechaCreacion: cursoV.FechaCreacion,
		Estado:        cursoV.Estado,
	}

	err := cursoClient.POSTcrearCurso(cursoData)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return cursoV, nil

}

func (s *cursoService) GetCursoPorID(id int) (*dto.CursosDatadto, *e.RestErr) {

	curso, err := cursoClient.GetCursoPorID(id)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	cursoDto := &dto.CursosDatadto{
		ID:                curso.ID,
		Nombre:            curso.Nombre,
		Descripcion:       curso.Descripcion,
		FechaCreacion:     curso.FechaCreacion,
		FechaModificacion: curso.FechaModificacion,
		Estado:            curso.Estado,
	}

	return cursoDto, nil

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

func (s *cursoService) PUTmodificarCurso(cursoV *dto.CursoClienteDataDto) (*dto.CursoClienteDataDto, *e.RestErr) {
	cursoData := &users.CursosData{
		ID:                cursoV.ID,
		Nombre:            cursoV.Nombre,
		Descripcion:       cursoV.Descripcion,
		FechaModificacion: cursoV.FechaModificacion,
		Estado:            cursoV.Estado,
	}

	_, CheckAdmin := s.IsAdmin(cursoV.User)

	if !CheckAdmin {
		return nil, &e.RestErr{Message: "No es admin, no puede modificar curso", StatusCode: 400}

	}

	err := cursoClient.PUTmodificarCurso(cursoData)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return cursoV, nil
}

func (s *cursoService) EliminarCurso(cursoV *dto.CursoClienteDataDto) *e.RestErr {

	cursoData := &users.CursosData{
		ID: cursoV.ID,
	}

	_, CheckAdmin := s.IsAdmin(cursoV.User)

	if !CheckAdmin {
		return &e.RestErr{Message: "No es admin, no puede eliminar curso", StatusCode: 400}

	}

	err := cursoClient.EliminarCurso(cursoData)
	if err != nil {
		return &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return nil
}
