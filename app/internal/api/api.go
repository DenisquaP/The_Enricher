package api

import (
	"enricher/database/models"
	"enricher/database/postgres"
	"enricher/internal/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var request entity.Request
	var response models.User

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	if err := ctx.BindJSON(&request); err != nil {
		log.Fatal(err)
	}

	response.Name = request.Name
	response.Surname = request.Surname
	response.Patronymic = request.Patronymic
	response.Age = 21
	response.Gender = "male"

	err = pg.InsertUser(response)
	if err != nil {
		log.Fatal(err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, response)
	}

}
