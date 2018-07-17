package main

import (
	"./config"
	"./web"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialise app config
	config.Init()

	router := gin.Default()

	// register app router
	web.Register(router)

	router.Run()
}
