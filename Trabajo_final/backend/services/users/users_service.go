package services

import (
	clientClients "backend/clients/client"
	"backend/domain/users"
	"backend/dto"
	e "backend/utils"
)

type clientService struct{}

type clientServiceInterface interface {
	POSTregistro(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr)
	POSTlogin(cliente *dto.ClientDatadto) (string, *e.RestErr)
}

var (
	ClientService clientServiceInterface
)

func init() {
	ClientService = &clientService{}
}

func (s *clientService) POSTregistro(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr) {
	client := &users.ClientData{
		FirstName: cliente.FirstName,
		LastName:  cliente.LastName,
		Email:     cliente.Email,
		User:      cliente.User,
		Password:  cliente.Password,
		Profesor:  cliente.Profesor,
	}

	err := clientClients.POSTregistro(client)
	if !err {
		return nil, &e.RestErr{StatusCode: 500}
	}

	return cliente, nil

}

func (s *clientService) POSTlogin(cliente *dto.ClientDatadto) (string, *e.RestErr) {

	client := &users.ClientData{
		User:     cliente.User,
		Email:    cliente.Email,
		Password: cliente.Password,
	}

	tokenString, err := clientClients.POSTlogin(client)

	if err != nil {
		return "", &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return tokenString, nil
}
