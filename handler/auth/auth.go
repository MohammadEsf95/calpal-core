package auth

import (
	"calpal-core/entity"
	"calpal-core/repository"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo repository.AuthRepository
}

func NewAuthHandler(r repository.AuthRepository) *Handler {
	return &Handler{repo: r}
}

func (a *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var signUpUser entity.User

	// Decode json body
	err := json.NewDecoder(r.Body).Decode(&signUpUser)
	if err != nil {
		handleErr(err, http.StatusBadRequest, w)
		return
	}

	// Validate sign up entity
	if err = ValidateUser(signUpUser); err != nil {
		handleErr(err, http.StatusBadRequest, w)
		return
	}

	uid, err := a.repo.SignUp(signUpUser)
	if err != nil {
		handleErr(err, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(uid)
}

func (a *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
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
