package users

import "github.com/jinzhu/gorm"

type ClientData struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(350) not null"`
	LastName  string `gorm:"type:varchar(350) not null"`
	Email     string `gorm:"type:varchar(350) not null; unique"`
	User      string `gorm:"type:varchar(350) not null; unique"`
	Password  string `gorm:"type:varchar(350) not null"`
	Admin     bool   `gorm:"not null"`

	Cursos []CursosData `gorm:"many2many:client_cursos;"`
}

type Clientes []ClientData
