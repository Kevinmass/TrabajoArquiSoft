package dto

// ForodataDto estructura de datos para el foro

type ForodataDto struct {
	CursoID uint   `json:"curso_id"`
	User    string `json:"user"`
	Comment string `json:"comment"`
}

type ForodataDtos []ForodataDto
