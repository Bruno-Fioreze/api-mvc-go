package routes

import (
	"github.com/Bruno-Fioreze/api-mvc-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:params", controller.FindUserByParams)
	r.DELETE("/user/:userId", controller.DeleteUser)
	r.POST("/user/", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
}
