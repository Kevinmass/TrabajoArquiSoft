package dto

type CursoClienteDataDto struct {
	User        string `json:"user"`
	ID          uint   `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Requisitos  string `json:"requisitos"`
	Duracion    string `json:"duracion"`
}
