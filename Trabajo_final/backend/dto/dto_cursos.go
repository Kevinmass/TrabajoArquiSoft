package dto

type CursosDatadto struct {
	ID          uint   `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	User        string `json:"user"`
	Duracion    string `json:"duracion"`
	Requisitos  string `json:"requisitos"`
}

type Cursosdto []CursosDatadto
