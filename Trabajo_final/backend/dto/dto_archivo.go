package dto

type ArchivodataDto struct {
	Base64 string `json:"base64" db:"base64"`
}

type ArchivoDto []ArchivodataDto
