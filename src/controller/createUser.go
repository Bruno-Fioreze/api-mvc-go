package controller

import (
	resterr "github.com/Bruno-Fioreze/api-mvc-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := resterr.NewBadRequestError("you called the route incorrectly")
	c.JSON(err.Code, err)
}
