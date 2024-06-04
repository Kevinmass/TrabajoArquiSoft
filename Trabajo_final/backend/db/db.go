package db

import (
	clientClient "backend/clients/client"
	cursoClient "backend/clients/cursos"
	inscripcionClient "backend/clients/inscripcion"

	"backend/domain/users"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "railway"                          //Nombre de la base de datos local de ustedes
	DBUser := "root"                             //usuario de la base de datos, habitualmente root
	DBPass := "qIBEPtXXvlWbTdwxEarjVqINFVOohogz" //password del root en la instalacion
	DBHost := "roundhouse.proxy.rlwy.net"        //host de la base de datos. hbitualmente 127.0.0.1
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":38957)/"+DBName+"?charset=utf8&parseTime=True")
	// @:/railway
	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	//migrar los clients

	clientClient.Db = db
	cursoClient.Db = db
	inscripcionClient.Db = db

}

func StartDB() {

	db.AutoMigrate(&users.ClientData{})
	db.AutoMigrate(&users.CursosData{})
	db.AutoMigrate(&users.InscriptosData{})

	log.Info("Database Migrated")
}
