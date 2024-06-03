package users

type InscriptosData struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	ClienteID uint `gorm:"not null"`
	CursoID   uint `gorm:"not null"`
}

type Inscriptos []InscriptosData
