package main

import (
	"enricher/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/create_user", api.CreateUser)

	router.Run()
}
