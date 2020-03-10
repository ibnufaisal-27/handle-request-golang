package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//Response Struct
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []string
}

func RespondSatu(w http.ResponseWriter, r *http.Request) {
	var response Response
	var debug_mode string
	if os.Getenv("APP_ENVIRONMENT") == "DEVELOPMENT" {
		env := os.Getenv("APP_ENVIRONMENT")
		port := os.Getenv("DEVELOPMENT_PORT")
		app_name := os.Getenv("DEVELOPMENT_APP_NAME")
		debug_mode = os.Getenv("DEVELOPMENT_DEBUG_MODE")

		response.Status = 200
		response.Message = "Success"
		response.Data = []string{"ENVIRONMENT :" + env, "app_port :" + port, "app_name :" + app_name, "debug_mode : " + debug_mode}

	} else if os.Getenv("APP_ENVIRONMENT") == "PRODUCTION" {
		env := os.Getenv("APP_ENVIRONMENT")
		port := os.Getenv("PRODUCTION_PORT")
		app_name := os.Getenv("PRODUCTION_APP_NAME")
		debug_mode = os.Getenv("PRODUCTION_DEBUG_MODE")

		response.Status = 200
		response.Message = "Success"
		response.Data = []string{"ENVIRONMENT :" + env, "app_port :" + port, "app_name :" + app_name, "debug_mode : " + debug_mode}
	}

	if debug_mode == "TRUE" {
		fmt.Println("http_status: ", response.Status)
		fmt.Println("message: ", response.Message)
		fmt.Println("data: ", response.Data)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func RespondDua(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("pong")
}
func RespondTiga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("waktu")
}
