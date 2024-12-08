package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-sqlite-backend/database"
	"go-sqlite-backend/services"
	"go-sqlite-backend/validators"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ValidateName(user.Name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ValidateEmail(user.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ValidatePassword(user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s added", user.Name)
}
