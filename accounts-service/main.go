package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/HashimovH/accounts-service/pkg/db"
	"github.com/HashimovH/accounts-service/pkg/handlers"
)

func main(){
	DB := db.Init()

	h := handlers.NewHandler(DB)
	router := mux.NewRouter()

	router.HandleFunc("/register", h.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/login", h.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/validate", h.ValidateUser).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":8080", router)

}