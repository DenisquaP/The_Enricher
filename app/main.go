package main

import (
	"enricher/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Db, err := postgres.NewPostgres()

	// if err != nil {
	// 	panic(err)
	// }

	router.POST("/create_user", api.CreateUser)

	router.Run()
}
