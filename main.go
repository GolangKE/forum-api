package main

import (
	"./web"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	web.Register(router)

	router.Run()
}
