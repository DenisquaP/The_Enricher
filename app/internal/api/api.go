package api

import (
	"enricher/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var request entity.Request
	var response entity.Response

	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	response.Request = request
	response.Age = 21
	response.Gender = "male"

	ctx.IndentedJSON(http.StatusOK, response)

}
