package adapters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HashimovH/accounts-service/pkg/models"
	"github.com/gorilla/mux"
)

type UserService interface {
	Create(user *models.User) (*models.User, error)
}

type HTTPAdapter struct {
	userService UserService
}

func UserHTTPAdapter(us UserService) *HTTPAdapter {
	return &HTTPAdapter{
		userService: us,
	}
}

func (h *HTTPAdapter) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.RegisterUser).Methods(http.MethodPost)
}

func (h *HTTPAdapter) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = r.Body.Close()
	if err != nil {
		log.Fatal("Could not close body")
	}

	result, err := h.userService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatalf("Could not encode json, err: %v", err)
	}
}

// func (h handler) ValidateUser(w http.ResponseWriter, r *http.Request) {}
// func (h handler) LoginUser(w http.ResponseWriter, r *http.Request)    {}
