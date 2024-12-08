package handlers

import (
	"go-sqlite-backend/services"
	"log"
	"net/http"
	"strconv"
)

func DeleteCountryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	err = services.DeleteCountry(countryID)
	if err != nil {
		http.Error(w, "Error deleting country", http.StatusInternalServerError)
		log.Println("Error deleting country:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Country deleted successfully"))
}
