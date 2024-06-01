package main

import (
	"backend/app"
	"backend/db"
)

func main() {

	db.StartDB()
	app.StartApplication()

}
