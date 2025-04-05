package handler

import (
	"calpal-core/contract"
	"calpal-core/entity"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	repo contract.AuthRepository
}

func NewAuthHandler(r contract.AuthRepository) *AuthHandler {
	return &AuthHandler{repo: r}
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var signUpUser entity.User

	err := json.NewDecoder(r.Body).Decode(&signUpUser)
	if err != nil {
		handleErr(err, http.StatusBadRequest, w)
		return
	}

	err = a.repo.SignUp(signUpUser)
	if err != nil {
		handleErr(err, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signUpUser)
}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func handleErr(err error, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	errorResponse := entity.ErrorResponse{Message: err.Error()}
	bytes, err := json.Marshal(errorResponse) // Handle error from Marshal
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
