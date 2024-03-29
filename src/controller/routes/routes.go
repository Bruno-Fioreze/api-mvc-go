package routes

import (
	"github.com/Bruno-Fioreze/api-mvc-go/src/controller"
	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.FindUserByID)
	r.GET("/user/:userEmail", controller.FindUserByEmail)
	r.DELETE("/user/:userId", controller.DeleteUser)
	r.POST("/user/", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
}
