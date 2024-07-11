package users

type CursosData struct {
	ID               uint   `gorm:"primary_key;auto_increment"`
	Nombre           string `gorm:"type:varchar(350); not null;case_sensitive:true"`
	ProfesorNombre   string `gorm:"type:varchar(350); not null"`
	ProfesorApellido string `gorm:"type:varchar(350); not null"`
	ProfesorCorreo   string `gorm:"type:varchar(350); not null"`
	Descripcion      string `gorm:"type:varchar(350); not null"`
	Duracion         string `gorm:"type:varchar(350); not null"`
	Requisitos       string `gorm:"type:varchar(350); not null"`
	FechaCreacion    string `gorm:"type:varchar(350) not null"`
}

type Cursos []CursosData
