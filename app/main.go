package main

import (
	"enricher/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/create_user", api.CreateUser)
	router.POST("/update_user", api.UpdateUser)

	router.Run()
}
