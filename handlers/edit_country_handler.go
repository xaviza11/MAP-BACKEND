package handlers

import (
	"encoding/json"
	"go-sqlite-backend/services"
	"net/http"
	"strconv"
)

func EditCountryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	countryIDStr := r.URL.Query().Get("country_id")
	if countryIDStr == "" {
		http.Error(w, "Country ID is required", http.StatusBadRequest)
		return
	}

	countryID, err := strconv.Atoi(countryIDStr)
	if err != nil {
		http.Error(w, "Invalid country ID", http.StatusBadRequest)
		return
	}

	var countryData struct {
		Name string  `json:"name"`
		LAT  float64 `json:"lat"`
		LON  float64 `json:"lon"`
	}

	err = json.NewDecoder(r.Body).Decode(&countryData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = services.EditCountry(countryID, countryData.Name, countryData.LAT, countryData.LON)
	if err != nil {
		http.Error(w, "Error editing country", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Country edited successfully"))
}
