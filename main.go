package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENVIRONMENT") == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			log.Println("Success loading .env for DEVELOPMENT")
		}
	} else if os.Getenv("APP_ENVIRONMENT") == "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			log.Println("Success loading .env for PRODUCTION")
		}
	}

	router := Route()

	fmt.Println("Connected to port 8000")
	http.ListenAndServe(":8000", router)

}
