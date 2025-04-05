package handler

import (
	"calpal-core/entity"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var signUpUser entity.User

	err := json.NewDecoder(r.Body).Decode(&signUpUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := entity.ErrorResponse{Message: err.Error()}
		bytes, _ := json.Marshal(errorResponse)
		w.Write(bytes)
		return
	}

	fmt.Printf("User %+v\n", signUpUser)
	// save to db
	//a.db.Query()
	json.NewEncoder(w).Encode(signUpUser)
}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign In")
}
