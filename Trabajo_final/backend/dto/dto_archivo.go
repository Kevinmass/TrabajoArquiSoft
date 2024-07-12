package dto

type ArchivodataDto struct {
	CursoNombre string `json:"curso_nombre"`
	Nombre      string `json:"nombre"`
	Ruta        string `json:"ruta"`
}

type ArchivoDto []ArchivodataDto
