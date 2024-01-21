package api

import (
	"enricher/database/models"
	"enricher/database/postgres"
	"enricher/internal/entity"
	"enricher/internal/services"
	"errors"
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
		ctx.JSON(http.StatusBadRequest, "can`t parse body")
		return
	}

	response.Name = request.Name
	response.Surname = request.Surname
	response.Patronymic = request.Patronymic

	age, err := services.Age(request.Name)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	response.Age = age

	gender, err := services.Gender(request.Name)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	response.Gender = gender

	nationality, err := services.Nationality(request.Name)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	response.Nationality = nationality

	err = pg.InsertUser(response)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)

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
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	err = pg.UpdateUser(request.FieldToUpdate, request.NewValue, request.UserId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	resp := entity.ResponseOk{Message: fmt.Sprintf("user %v has been updated by value %v: %v", request.UserId, request.FieldToUpdate, request.NewValue)}
	ctx.IndentedJSON(http.StatusOK, resp)
}

func DeleteUser(ctx *gin.Context) {
	var request entity.DeleteRequest

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
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	err = pg.DeleteUser(request.UserID)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	resp := entity.ResponseOk{
		Message: fmt.Sprintf("user %v has been deleted", request.UserID),
	}
	ctx.IndentedJSON(http.StatusOK, resp)
}

func GetUsers(ctx *gin.Context) {
	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	res, err := pg.GetUsers()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func GetUsersFilter(ctx *gin.Context) {
	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	filterTag := ctx.Query("filter_tag")
	if filterTag == "" {
		err := errors.New("got empry filter")
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	filter := ctx.Query("filter")
	if filter == "" {
		err := errors.New("got empry filter")
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	res, err := pg.GetUsersByFilter(filterTag, filter)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, res)

}
