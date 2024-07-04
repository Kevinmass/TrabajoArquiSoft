package users

type Forodata struct {
	ID         uint   `gorm:"primary_key;auto_increment"`
	CursoID    uint   `gorm:"type:int; not null"`
	User       string `gorm:"type:varchar(350); not null"`
	Comentario string `gorm:"type:varchar(350); not null"`
}
