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

func GuardarArchivo(file *dto.ArchivodataDto, cursoId string) *e.RestErr {
	archivo := &users.Archivodata{
		Base64:  file.Base64,
		CursoId: cursoId,
	}
	log.Println(archivo.Base64)

	result := Db.Create(archivo)
	if err := result.Error; err != nil {
		return &e.RestErr{Message: "Cannot save to database", StatusCode: 500}
	}

	return nil
}

func GetFiles(cursoId string) ([]dto.ArchivodataDto, *e.RestErr) {
	var archivos []users.Archivodata
	if err := Db.Where("curso_id = ?", cursoId).Find(&archivos).Error; err != nil {
		return nil, &e.RestErr{Message: "Cannot fetch files from database", StatusCode: 500}
	}

	var archivosDto []dto.ArchivodataDto
	for _, archivo := range archivos {
		archivosDto = append(archivosDto, dto.ArchivodataDto{
			Base64: archivo.Base64,
		})
	}

	return archivosDto, nil
}
