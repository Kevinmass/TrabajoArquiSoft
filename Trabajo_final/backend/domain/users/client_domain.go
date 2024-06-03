package users

type ClientData struct {
	ID                uint   `gorm:"primary_key;auto_increment"`
	FechaCreacion     string `gorm:"type:varchar(350) not null"`
	FechaEliminacion  string `gorm:"type:varchar(350) not null"`
	FechaModificacion string `gorm:"type:varchar(350) not null"`

	FirstName string `gorm:"type:varchar(350) not null"`
	LastName  string `gorm:"type:varchar(350) not null"`
	Email     string `gorm:"type:varchar(350) not null; unique"`
	User      string `gorm:"type:varchar(350) not null; unique"`
	Password  string `gorm:"type:varchar(350) not null"`
	Admin     bool   `gorm:"not null"`
}

type Clientes []ClientData
