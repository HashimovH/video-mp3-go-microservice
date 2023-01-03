package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HashimovH/accounts-service/pkg/models"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
func (h handler) ValidateUser(w http.ResponseWriter, r *http.Request) {}
func (h handler) LoginUser(w http.ResponseWriter, r *http.Request)    {}
