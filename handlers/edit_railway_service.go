package handlers

import (
	"encoding/json"
	"go-sqlite-backend/database"
	"go-sqlite-backend/services"
	"net/http"
	"strconv"
)

func EditRailwaysHandler(w http.ResponseWriter, r *http.Request) {
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

	var railways []struct {
		Name string `json:"name"`
	}

	err = json.NewDecoder(r.Body).Decode(&railways)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var railwayList []database.Railway
	for _, railway := range railways {
		railwayList = append(railwayList, database.Railway{
			Name: railway.Name,
		})
	}

	err = services.EditRailways(countryID, railwayList)
	if err != nil {
		http.Error(w, "Error editing railways", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Railways edited successfully"))
}
