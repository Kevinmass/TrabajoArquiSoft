package users

type Archivodata struct {
	ID      uint   `gorm:"primary_key;auto_increment"`
	Base64  string `gorm:"type:mediumtext"`
	CursoId string `gorm:"type:varchar(100)"`
}
type Archivo []Archivodata
