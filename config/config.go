package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// use gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var repo *gorm.DB

// Init config
func Init() *gorm.DB {
	// start db connection
	dialect := os.Getenv("DB_DIALECT")
	db, err := gorm.Open(dialect, dbConfig())

	if err != nil {
		panic("failed to connect database")
	}

	// set logMode to true
	db.LogMode(true)

	// assign gorm.DB instance to repo
	repo = db

	return repo
}

// Repo returns database
func Repo() *gorm.DB {
	return repo
}

// unexported functions

func dbConfig() string {
	str := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USERNAME") +
		" dbname=" + os.Getenv("DATABASE") +
		" password=" + os.Getenv("DB_PASSWORD")

	if gin.IsDebugging() {
		// disable ssl mode if we are in development
		str = str + " sslmode=disable"
	}

	return str
}
