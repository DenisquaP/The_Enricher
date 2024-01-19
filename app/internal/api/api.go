package api

import (
	"enricher/database/models"
	"enricher/database/postgres"
	"enricher/internal/entity"
	"enricher/internal/services"
	"fmt"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var request entity.CreateRequest
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

	age, err := services.Age(request.Name)
	if err != nil {
		log.Fatal(err)
	}

	response.Age = age

	gender, err := services.Gender(request.Name)
	if err != nil {
		log.Fatal(err)
	}

	response.Gender = gender

	nationality, err := services.Nationality(request.Name)
	if err != nil {
		log.Fatal(err)
	}

	response.Nationality = nationality

	err = pg.InsertUser(response)
	if err != nil {
		log.Fatal(err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, response)
	}
}

func UpdateUser(ctx *gin.Context) {
	var request entity.UpdateRequest

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

	err = pg.UpdateUser(request.FieldToUpdate, request.NewValue, request.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprint(err)})
	} else {
		ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
	}

}
