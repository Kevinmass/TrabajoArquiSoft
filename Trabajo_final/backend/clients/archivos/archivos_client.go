package database

import (
	"backend/domain/users"
	"backend/dto"
	e "backend/utils"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func SaveFile(file *dto.ArchivodataDto) *e.RestErr {
	archivo := &users.Archivodata{
		Base64: file.Base64,
	}
	log.Println(archivo.Base64)

	// Use GORM's Create method for safe insertion
	result := Db.Create(archivo)
	if err := result.Error; err != nil {
		return &e.RestErr{Message: "Cannot save to database", StatusCode: 500}
	}

	return nil
}
