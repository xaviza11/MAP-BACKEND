package handlers

import (
	"encoding/json"
	"go-sqlite-backend/database"
	"go-sqlite-backend/services"
	"net/http"
	"strconv"
)

func EditFactionsHandler(w http.ResponseWriter, r *http.Request) {
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

	var factions []struct {
		Name    string `json:"name"`
		Support string `json:"support"`
	}

	err = json.NewDecoder(r.Body).Decode(&factions)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var factionList []database.Faction
	for _, faction := range factions {
		factionList = append(factionList, database.Faction{
			Name:    faction.Name,
			Support: faction.Support,
		})
	}

	err = services.EditFactions(countryID, factionList)
	if err != nil {
		http.Error(w, "Error editing factions", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Factions edited successfully"))
}
