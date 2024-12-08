package handlers

import (
	"encoding/json"
	"go-sqlite-backend/database"
	"go-sqlite-backend/services"
	"log"
	"net/http"
)

func CreateCountryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Country     database.Country     `json:"country"`
		CountryInfo database.CountryInfo `json:"country_info"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Println("Error decoding request body:", err)
		return
	}

	err = services.CreateCountry(requestBody.Country, requestBody.CountryInfo)
	if err != nil {
		http.Error(w, "Error creating country", http.StatusInternalServerError)
		log.Println("Error creating country:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Country created successfully"))
}
