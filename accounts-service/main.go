package main

import (
	"log"
	"net/http"

	"github.com/HashimovH/accounts-service/pkg/db"
	"github.com/HashimovH/accounts-service/pkg/handlers"
	"github.com/HashimovH/accounts-service/pkg/repositories"
	"github.com/HashimovH/accounts-service/pkg/services"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	DB := db.Init()
	userRepository := repositories.NewRepository(DB)
	userService := services.NewService(userRepository)
	userHTTPHandler := handlers.NewHandler(userService)
	router := mux.NewRouter()
	userHTTPHandler.RegisterRoutes(router)

	router.HandleFunc("/register", h.RegisterUser).Methods(http.MethodPost)
	// router.HandleFunc("/login", h.LoginUser).Methods(http.MethodPost)
	// router.HandleFunc("/validate", h.ValidateUser).Methods(http.MethodPost)
	return router

}

func main() {
	router := Router()
	log.Println("API is running!")
	http.ListenAndServe(":8080", router)

}
