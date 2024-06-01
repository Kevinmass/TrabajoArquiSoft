package dto

type InscriptosDataDto struct {
	ClienteID       uint   `json:"cliente_id"`
	CursoID         uint   `json:"curso_id"`
	ID              int    `json:"id"`
	InscriptionDate string `json:"fecha_inscripcion"`
}
