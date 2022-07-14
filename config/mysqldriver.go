package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	errENV := godotenv.Load()
	if errENV != nil {
		log.Fatalf("error loading env file")
	}
	config := map[string]string{
		"DB_Username": os.Getenv("DBusername"),
		"DB_Password": os.Getenv("DBpassword"),
		"DB_Port":     os.Getenv("DBport"),
		"DB_Host":     os.Getenv("DBhost"),
		"DB_Name":     os.Getenv("DBname"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	return db
}
