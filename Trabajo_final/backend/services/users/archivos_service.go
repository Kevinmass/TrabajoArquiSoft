package services

import (
	"encoding/base64"
	"io"
	"mime/multipart"

	archivoClient "backend/clients/archivos"
	"backend/dto"
	e "backend/utils"
)

type fileUploadService struct{}

type fileUploadServiceInterface interface {
	SubirArchivos(fileHeader *multipart.FileHeader, cursoId string) *e.RestErr
}

var (
	FileUploadService fileUploadServiceInterface = &fileUploadService{}
)

func (s *fileUploadService) SubirArchivos(fileHeader *multipart.FileHeader, cursoId string) *e.RestErr {

	fileContent, err := fileHeader.Open()
	if err != nil {
		return &e.RestErr{Message: "Cannot open the file", StatusCode: 500}
	}
	defer fileContent.Close()

	data, err := io.ReadAll(fileContent)
	if err != nil {
		return &e.RestErr{Message: "Cannot read the file", StatusCode: 500}
	}

	base64String := base64.StdEncoding.EncodeToString(data)
	archivo := dto.ArchivodataDto{Base64: base64String}

	if err := archivoClient.GuardarArchivo(&archivo, cursoId); err != nil {
		return &e.RestErr{Message: "Cannot save to database", StatusCode: 500}
	}

	return nil
}

type fileService struct{}

type fileServiceInterface interface {
	GetArchivos(string) ([]dto.ArchivodataDto, *e.RestErr)
}

var (
	FileService fileServiceInterface = &fileService{}
)

func (s *fileService) GetArchivos(cursoId string) ([]dto.ArchivodataDto, *e.RestErr) {
	return archivoClient.GetFiles(cursoId)
}
