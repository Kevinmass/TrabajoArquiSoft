package users

type CursosData struct {
	ID                int    `gorm:"primary_key;auto_increment; not null"`
	Nombre            string `gorm:"type:varchar(350); not null"`
	Descripcion       string `gorm:"type:varchar(350); not null"`
	FechaCreacion     string `gorm:"type:varchar(350); not null"`
	FechaEliminacion  string `gorm:"type:varchar(350); not null"`
	FechaModificacion string `gorm:"type:varchar(350); not null"`
	Estado            bool   `gorm:"type:boolean; not null"`
}

type Cursos []CursosData
