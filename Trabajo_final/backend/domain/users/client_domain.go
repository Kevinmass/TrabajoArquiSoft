package users

type ClientData struct {
	ID            uint   `gorm:"primary_key;auto_increment"`
	FirstName     string `gorm:"type:varchar(350) not null"`
	LastName      string `gorm:"type:varchar(350) not null"`
	Email         string `gorm:"type:varchar(350) not null; unique; case_sensitive:true"`
	User          string `gorm:"type:varchar(350) not null; unique; case_sensitive:true"`
	Password      string `gorm:"type:varchar(350) not null"`
	Profesor      bool   `gorm:"not null"`
	FechaCreacion string `gorm:"type:varchar(350) not null"`
}

type Clientes []ClientData
