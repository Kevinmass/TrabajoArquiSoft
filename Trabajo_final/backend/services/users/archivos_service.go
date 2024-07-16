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
	UploadFile(fileHeader *multipart.FileHeader) *e.RestErr
}

var (
	FileUploadService fileUploadServiceInterface
)

func init() {
	FileUploadService = &fileUploadService{}
}

func (s *fileUploadService) UploadFile(fileHeader *multipart.FileHeader) *e.RestErr {

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

	if err := archivoClient.SaveFile(&archivo); err != nil {
		return &e.RestErr{Message: "Cannot save to database", StatusCode: 500}
	}

	return nil
}
