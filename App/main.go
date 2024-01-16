package main

import (
	api "enricher/internal/API"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/create_user", api.CreateUser(c*gin.Context))

	r.Run()
}
