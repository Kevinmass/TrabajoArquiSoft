package users

import (
	"time"
)

type InscriptosData struct {
	ClienteID       uint      `gorm:"not null"`
	CursoID         uint      `gorm:"not null"`
	InscriptionDate time.Time `gorm:"not null"`
}

type Inscriptos []InscriptosData
