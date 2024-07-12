package users

type CursosData struct {
	ID          uint   `gorm:"primary_key;"`
	Nombre      string `gorm:"type:varchar(350); not null;case_sensitive:true; unique"`
	User        string `gorm:"type:varchar(350); not null;case_sensitive:true"`
	Descripcion string `gorm:"type:varchar(350); not null"`
	Duracion    string `gorm:"type:varchar(350); not null"`
	Requisitos  string `gorm:"type:varchar(350); not null"`
}

type Cursos []CursosData
