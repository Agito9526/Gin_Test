package main

import (
	. "GolangAPI/controller"
	"GolangAPI/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	go func() {
		database.DD()
	}()

	router.Run(":8080")
}
