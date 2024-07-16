package users

type Archivodata struct {
	ID     uint   `gorm:"primary_key;auto_increment"`
	Base64 string `gorm:"type:mediumtext"`
}
type Archivo []Archivodata
