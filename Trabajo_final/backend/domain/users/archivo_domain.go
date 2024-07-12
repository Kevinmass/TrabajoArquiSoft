package users

type Archivodata struct {
	ID          uint   `gorm:"primary_key;auto_increment"`
	CursoNombre string `gorm:"varchar(350); not null"`
	Nombre      string `gorm:"type:varchar(350); not null"`
	Ruta        string `gorm:"type:varchar(350); not null"`
}
