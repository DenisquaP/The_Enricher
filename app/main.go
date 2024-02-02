package main

import (
	"enricher/database/postgres"
	"enricher/internal/api"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.MigrationsUp()
	if err := pg.MigrationsUp(); err.Error() != "no change" {
		log.Fatalf("Migration create was failed: %v", err)
	}

	router.POST("/create_user", api.CreateUser)
	router.PUT("/update_user", api.UpdateUser)
	router.DELETE("/delete_user", api.DeleteUser)

	router.GET("/get_users", api.GetUsers)
	router.GET("/get_users_by_filter", api.GetUsersFilter)

	router.Run()
}
