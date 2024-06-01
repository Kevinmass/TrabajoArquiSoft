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
	PUThashearPassword(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr)
	GETverificarCliente(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr)
	GETloginCliente(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr)
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
		Admin:     cliente.Admin,
	}

	err := clientClients.POSTregistro(client)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return cliente, nil

}

func (s *clientService) PUThashearPassword(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr) {

	client := &users.ClientData{
		Password: cliente.Password,
	}

	err := clientClients.PUThashearPassword(client)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return cliente, nil

}

func (s *clientService) GETverificarCliente(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr) {

	client := &users.ClientData{
		Email:    cliente.Email,
		Password: cliente.Password,
	}

	err := clientClients.GETverificarCliente(client)
	if err != nil {
		return nil, &e.RestErr{Message: err.Error(), StatusCode: 500}
	}

	return cliente, nil
}

func (s *clientService) GETloginCliente(cliente *dto.ClientDatadto) (*dto.ClientDatadto, *e.RestErr) {

	client := &users.ClientData{
		Email:    cliente.Email,
		Password: cliente.Password,
	}

	clientClients.GETloginCliente(client)

	//verificar si el mail y la contraseña coincide con lo que regresa la funcion GETloginCliente

	return cliente, nil
}
