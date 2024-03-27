package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, world!")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
