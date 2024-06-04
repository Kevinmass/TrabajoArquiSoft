package dto

type CursosDatadto struct {
	ID                uint   `json:"id"`
	Nombre            string `json:"nombre"`
	Descripcion       string `json:"descripcion"`
	ProfesorNombre    string `json:"profesor_nombre"`
	ProfesorApellido  string `json:"profesor_apellido"`
	ProfesorCorreo    string `json:"profesor_correo"`
	FechaCreacion     string `json:"fecha_creacion"`
	FechaModificacion string `json:"fecha_actualizacion"`
}

type Cursosdto []CursosDatadto