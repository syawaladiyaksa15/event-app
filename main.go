package main

import (
	"project/group3/config"
	"project/group3/factory"
	"project/group3/migration"
	"project/group3/routes"
)

func main() {
	//initiate db connection
	dbConn := config.InitDB()

	// run auto migrate from gorm
	migration.InitMigrate(dbConn)

	// initiate factory
	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)

	e.Start(":8000")

}
