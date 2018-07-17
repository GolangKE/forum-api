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
func Init() {
	// start db connection
	initDB()
}

// Repo returns database
func Repo() *gorm.DB {
	return repo
}

// unexported functions

// InitDB initialises database connection
func initDB() {
	db, err := gorm.Open(os.Getenv("DB_DIALECT"), dbConfig())

	if err != nil {
		panic("failed to connect database")
	}

	repo = db
	defer db.Close()
}

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
