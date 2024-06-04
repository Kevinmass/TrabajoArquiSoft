package client

import (
	"backend/domain/users"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	Db *gorm.DB
)

func POSTregistro(cliente *users.ClientData) error {

	//se verifica si el cliente ya existe en la base de datos
	var clienteDB users.ClientData
	err := Db.Where("email = ?", cliente.Email).First(&clienteDB).Error
	if err == nil {
		log.Error("Cliente already exists")
		return err
	}

	err = Db.Where("user = ?", cliente.User).First(&clienteDB).Error
	if err == nil {
		log.Error("Cliente already exists")
		return err
	}

	//se encripta la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(cliente.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error hashing password: ", err)
		return err
	}
	cliente.Password = string(hash)

	//se crea el cliente
	err = Db.Create(&cliente).Error
	if err != nil {
		log.Error("Error creating cliente: ", err)
		return err
	}

	return nil

}

// verficar si el cliente ya existe en la base de datos, y evitar que se cree un cliente con el mismo email y usuario

func POSTlogin(cliente *users.ClientData) bool {
	var clienteDB users.ClientData
	err := Db.Where("email = ?", cliente.Email).First(&clienteDB).Error
	if err != nil {
		log.Error("Error getting cliente: ", err)
		return false
	}

	err = Db.Where("user = ?", cliente.User).First(&clienteDB).Error
	if err != nil {
		log.Error("Error getting cliente: ", err)
		return false
	}

	// se compara la contraseña encriptada con la contraseña ingresada
	err = bcrypt.CompareHashAndPassword([]byte(clienteDB.Password), []byte(cliente.Password))
	if err != nil {
		log.Error("Error comparing password: ", err)
		return false
	}

	// generar token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  cliente.User,
		"email": cliente.Email,
	})

	tokenString, err := token.SignedString([]byte("asdhsghghsjgasgjgfd2737187ga"))
	if err != nil {
		log.Error("Error generating token string: ", err)
		return false
	}

	log.Debug("Token: ", tokenString)

	// verificar token de cliente

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("asdhsghghsjgasgjgfd2737187ga"), nil
	})

	//encontrar el cliente por el token

	if err != nil {
		log.Error("Error parsing token: ", err)
		return false
	}

	if token.Valid {
		//buscar el cliente por el token
		err = Db.Where("user = ?", token.Claims.(jwt.MapClaims)["user"]).First(&clienteDB).Error
		if err != nil {
			log.Error("Error getting cliente: ", err)
			return false
		}
		return true

	} else {
		return false
	}

}

func GETclientevalidado(cliente *users.ClientData) (*users.ClientData, error) {

	if POSTlogin(cliente) == true {
		return cliente, nil
	} else {
		return nil, nil
	}
}
