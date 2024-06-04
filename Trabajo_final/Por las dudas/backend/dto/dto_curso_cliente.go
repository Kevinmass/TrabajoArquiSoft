package dto

type CursoClienteDataDto struct {
	User              string `json:"user"`
	ID                uint   `json:"id"`
	Nombre            string `json:"nombre"`
	Descripcion       string `json:"descripcion"`
	NombreUsuario     string `json:"nombre_usuario"`
	ApellidoUsuario   string `json:"apellido_usuario"`
	CorreoUsuario     string `json:"correo_usuario"`
	FechaCreacion     string `json:"fecha_creacion"`
	FechaModificacion string `json:"fecha_actualizacion"`
}
