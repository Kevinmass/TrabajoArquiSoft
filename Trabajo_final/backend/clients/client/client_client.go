package client

import (
	"backend/domain/users"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var Db *gorm.DB

func POSTregistro(cliente *users.ClientData) error {
	err := Db.Create(cliente).Error
	if err != nil {
		log.Error("Error creating cliente: ", err)
		return err
	}
	return nil

}

// Modifica la contraseña en la base de datos y la encripta

func PUThashearPassword(cliente *users.ClientData) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cliente.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error hashing password: ", err)
		return err
	}
	cliente.Password = string(hashedPassword)
	err = Db.Save(cliente).Error
	if err != nil {
		log.Error("Error saving cliente: ", err)
		return err
	}
	return nil
}

// verficar si el cliente ya existe en la base de datos, y evitar que se cree un cliente con el mismo email y usuario

func GETverificarCliente(cliente *users.ClientData) error {
	var clienteDB users.ClientData
	err := Db.Where("email = ?", cliente.Email).First(&clienteDB).Error
	if err != nil {
		log.Error("Error getting cliente: ", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(clienteDB.Password), []byte(cliente.Password))
	if err != nil {
		log.Error("Error comparing passwords: ", err)
		return err
	}

	return nil

}

// loginCliente: verifica si el cliente existe en la base de datos y si la contraseña es correcta, login por token

func GETloginCliente(cliente *users.ClientData) users.ClientData {

	if err := GETverificarCliente(cliente); err != nil {
		GETverificarCliente(cliente)
	}
	return *cliente

}
