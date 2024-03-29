package main

import (
	"fmt"
	"log"

	"github.com/Bruno-Fioreze/api-mvc-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, world!")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	router.Run(":8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
