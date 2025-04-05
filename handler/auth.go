package handler

import (
	"calpal-core/entity"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
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
		handleErr(err, http.StatusBadRequest, w)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(signUpUser.Password), bcrypt.DefaultCost)
	if err != nil {
		handleErr(err, http.StatusBadRequest, w)
		return
	}

	signUpUser.Password = string(hashPassword)

	signUpUser.ID = uuid.New().String()

	// save to db
	_, err = a.db.Query("INSERT INTO users (id, first_name, last_name, password, email, created_at, updated_at)"+
		" VALUES ($1, $2, $3, $4, $5, $6, $7)", signUpUser.ID, signUpUser.FirstName, signUpUser.LastName, signUpUser.Password,
		signUpUser.Email, time.Now(), time.Now())
	if err != nil {
		handleErr(err, http.StatusInternalServerError, w)
		return
	}

	json.NewEncoder(w).Encode(signUpUser)
}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign In")
}

func handleErr(err error, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	errorResponse := entity.ErrorResponse{Message: err.Error()}
	bytes, _ := json.Marshal(errorResponse)
	w.Write(bytes)
}
