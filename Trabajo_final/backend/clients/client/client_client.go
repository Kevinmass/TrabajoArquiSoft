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

func POSTregistro(cliente *users.ClientData) bool {

	//se verifica si el cliente ya existe en la base de datos
	var clienteDB users.ClientData
	err := Db.Where("email = ?", cliente.Email).First(&clienteDB).Error
	if err == nil {
		log.Error("Cliente already exists")
		return false
	}

	err = Db.Where("user = ?", cliente.User).First(&clienteDB).Error
	if err == nil {
		log.Error("Cliente already exists")
		return false
	}

	//se encripta la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(cliente.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error hashing password: ", err)
		return false
	}
	cliente.Password = string(hash)

	//se crea el cliente
	err = Db.Create(&cliente).Error
	if err != nil {
		log.Error("Error creating cliente: ", err)
		return false
	}

	return true

}

// verficar si el cliente ya existe en la base de datos, y evitar que se cree un cliente con el mismo email y usuario

func POSTlogin(cliente *users.ClientData) (string, error) {
	var clienteDB users.ClientData
	var err error

	err = Db.Where("user = ?", cliente.User).First(&clienteDB).Error
	if err != nil {
		log.Error("Error getting cliente: ", err)
		return "", err
	}

	// se compara la contraseña encriptada con la contraseña ingresada
	err = bcrypt.CompareHashAndPassword([]byte(clienteDB.Password), []byte(cliente.Password))
	if err != nil {
		log.Error("Error comparing password: ", err)
		return "", err
	}

	// generar token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  cliente.User,
		"email": cliente.Email,
	})

	tokenString, err := token.SignedString([]byte("asdhsghghsjgasgjgfd2737187ga"))
	if err != nil {
		log.Error("Error generating token string: ", err)
		return "", err
	}

	log.Debug("Token: ", tokenString)

	// verificar token de cliente

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("asdhsghghsjgasgjgfd2737187ga"), nil
	})

	//encontrar el cliente por el token

	if err != nil {
		log.Error("Error parsing token: ", err)
		return "", err
	}

	if token.Valid {
		//buscar el cliente por el token
		err = Db.Where("user = ?", token.Claims.(jwt.MapClaims)["user"]).First(&clienteDB).Error
		if err != nil {
			log.Error("Error getting cliente: ", err)
			return "", err
		}
		return tokenString, nil

	} else {
		return "", err
	}

}
