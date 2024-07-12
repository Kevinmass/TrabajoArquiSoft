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
	IsAdmin(cliente *dto.CursoClienteDataDto) (*e.RestErr, bool)
	GetCursoPorID(id uint) (*dto.CursosDatadto, *e.RestErr)
	GetCursosTotales() ([]dto.CursosDatadto, *e.RestErr)
	PUTmodificarCurso(cursoV *dto.CursoClienteDataDto) (*dto.CursoClienteDataDto, *e.RestErr)
	EliminarCurso(cursoV *dto.CursoClienteDataDto) *e.RestErr
	GetCursosPorNombre(nombre string) ([]dto.CursosDatadto, *e.RestErr)
	GetTotalCursos() (int, *e.RestErr)
}

var (
	CursoService cursoServiceInterface
)

func init() {
	CursoService = &cursoService{}
}

// si el cliente es admin crear curso

func (s *cursoService) IsAdmin(cliente *dto.CursoClienteDataDto) (*e.RestErr, bool) {

	clienteData := &users.ClientData{
		User: cliente.User,
	}

	CheckAdmin := cursoClient.IsAdmin(clienteData)

	return nil, CheckAdmin
}

func (s *cursoService) POSTcrearCurso(cursoV *dto.CursoClienteDataDto) (*dto.CursoClienteDataDto, *e.RestErr) {

	cursoData := &users.CursosData{

		Nombre:      cursoV.Nombre,
		Descripcion: cursoV.Descripcion,
		User:        cursoV.User,
		Duracion:    cursoV.Duracion,
		Requisitos:  cursoV.Requisitos,
	}

	_, CheckAdmin := s.IsAdmin(cursoV)

	if CheckAdmin == false {
		return nil, &e.RestErr{Message: "No es admin, no puede crear curso", StatusCode: 400}

	}

	err := cursoClient.POSTcrearCurso(cursoData)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	cursoV.ID = cursoData.ID

	return cursoV, nil

}

func (s *cursoService) GetCursoPorID(id uint) (*dto.CursosDatadto, *e.RestErr) {

	curso, err := cursoClient.GetCursoPorID(id)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	cursoDto := &dto.CursosDatadto{
		ID:          curso.ID,
		Nombre:      curso.Nombre,
		Descripcion: curso.Descripcion,
		User:        curso.User,
		Duracion:    curso.Duracion,
		Requisitos:  curso.Requisitos,
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
			ID:          curso.ID,
			Nombre:      curso.Nombre,
			Descripcion: curso.Descripcion,
			User:        curso.User,
			Duracion:    curso.Duracion,
			Requisitos:  curso.Requisitos,
		}
		cursosDto = append(cursosDto, *cursoDto)
	}

	return cursosDto, nil
}

func (s *cursoService) PUTmodificarCurso(cursoV *dto.CursoClienteDataDto) (*dto.CursoClienteDataDto, *e.RestErr) {
	cursoData := &users.CursosData{

		ID:          cursoV.ID,
		Nombre:      cursoV.Nombre,
		Descripcion: cursoV.Descripcion,
		User:        cursoV.User,
		Duracion:    cursoV.Descripcion,
		Requisitos:  cursoV.Requisitos,
	}

	_, CheckAdmin := s.IsAdmin(cursoV)

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

	_, CheckAdmin := s.IsAdmin(cursoV)

	if !CheckAdmin {
		return &e.RestErr{Message: "No es admin, no puede eliminar curso", StatusCode: 400}

	}

	err := cursoClient.EliminarCurso(cursoData)
	if err != nil {
		return &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return nil
}

func (s *cursoService) GetCursosPorNombre(nombre string) ([]dto.CursosDatadto, *e.RestErr) {

	curso, err := cursoClient.GetCursosPorNombre(nombre)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	var err2 *e.RestErr
	err2 = nil

	var cursoDto []dto.CursosDatadto

	for _, curso := range curso {
		cursoDto = append(cursoDto, dto.CursosDatadto{
			ID:          curso.ID,
			Nombre:      curso.Nombre,
			Descripcion: curso.Descripcion,
			Duracion:    curso.Duracion,
			Requisitos:  curso.Requisitos,
		})
	}

	// si no se encuentra el curso con la palabra, envia una lista vacia
	if len(cursoDto) == 0 {
		err2 = &e.RestErr{StatusCode: 404}
		cursoDto = []dto.CursosDatadto{}

	}

	return cursoDto, err2

}

func (s *cursoService) GetTotalCursos() (int, *e.RestErr) {
	cursos, err := cursoClient.GetCursosTotales()
	if err != nil {
		return 0, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}
	return len(cursos), nil
}
