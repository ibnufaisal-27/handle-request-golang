package main

import (
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", RespondSatu).Methods("GET")
	router.HandleFunc("/ping", RespondDua).Methods("GET")
	router.HandleFunc("/time", RespondTiga).Methods("GET")

	return router
}
