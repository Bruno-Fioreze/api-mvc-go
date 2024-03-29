package controller

import (
	"fmt"

	resterr "github.com/Bruno-Fioreze/api-mvc-go/src/configuration/rest_err"
	"github.com/Bruno-Fioreze/api-mvc-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		message := fmt.Sprintf("There are some incorrect parameters, error=%s", err.Error())
		restErr := resterr.NewBadRequestError(message)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
