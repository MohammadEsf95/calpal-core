package user

import (
	"calpal-core/repository"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) Handler {
	return Handler{repo: repo}
}

func (h *Handler) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.repo.Users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
