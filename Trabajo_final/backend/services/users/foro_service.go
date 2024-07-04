package services

import (
	foroClient "backend/clients/foro"

	"backend/dto"
	e "backend/utils"
)

type foroService struct{}

type foroServiceInterface interface {
	POSTcomentario(forodata *dto.ForodataDto) *e.RestErr
	GETcomentarios(forodata *dto.ForodataDto) ([]dto.ForodataDto, *e.RestErr)
	DELETEcomentario(forodata *dto.ForodataDto) *e.RestErr
}

var (
	ForoService foroServiceInterface
)

func init() {

	ForoService = &foroService{}

}

func (s *foroService) POSTcomentario(forodata *dto.ForodataDto) *e.RestErr {

	err := foroClient.POSTcomentario(forodata.CursoID, forodata.User, forodata.Comment)

	if err.StatusCode != 200 {
		return &e.RestErr{Message: "Error al crear comentario", StatusCode: err.StatusCode}
	}

	return nil
}

func (s *foroService) GETcomentarios(forodata *dto.ForodataDto) ([]dto.ForodataDto, *e.RestErr) {

	comentarios, err := foroClient.GETcomentarios(forodata.CursoID)
	var err2 *e.RestErr = nil

	if err.StatusCode != 200 {
		return nil, &e.RestErr{Message: "Error al obtener comentarios", StatusCode: err.StatusCode}
	}

	var comentariosDto []dto.ForodataDto

	for _, comentario := range comentarios {
		comentarioDto := &dto.ForodataDto{
			CursoID: comentario.CursoID,
			User:    comentario.User,
			Comment: comentario.Comentario,
		}
		comentariosDto = append(comentariosDto, *comentarioDto)
	}

	if len(comentariosDto) == 0 {

		err2 = &e.RestErr{StatusCode: 404}
		comentariosDto = []dto.ForodataDto{}

	}

	return comentariosDto, err2
}

func (s *foroService) DELETEcomentario(forodata *dto.ForodataDto) *e.RestErr {

	err := foroClient.DELETEcomentario(forodata.User)

	if err.StatusCode != 200 {
		return &e.RestErr{Message: "Error al eliminar comentario", StatusCode: err.StatusCode}
	}

	return nil
}
