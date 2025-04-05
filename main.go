package main

import (
	"calpal-core/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Message: err.Error()}
		bytes, _ := json.Marshal(errorResponse)
		w.Write(bytes)
		return
	}

	fmt.Printf("User %+v\n", user)
	// save to db

	json.NewEncoder(w).Encode(user)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign In")
}

func setTargetCalories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented")
}

func main() {
	// User registration: signup / login
	http.HandleFunc("/sign-up", signUp)
	http.HandleFunc("/sign-in", signIn)

	// Set target calories
	http.HandleFunc("/set-target-calories", setTargetCalories)

	http.ListenAndServe(":8080", nil)
	// food log
}
