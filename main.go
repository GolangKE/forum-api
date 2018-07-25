package main

import (
	"./config"
	"./db"
	"./web"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialise app config
	repo := config.Init()

	// add seed data to database
	db.Seed(repo)

	defer repo.Close()

	// register app router
	router := gin.New()

	web.Register(router)

	// run application
	router.Run()
}
