package handlers

import (
	"encoding/json"
	"go-sqlite-backend/database"
	"go-sqlite-backend/services"
	"net/http"
	"strconv"
)

func EditExportsHandler(w http.ResponseWriter, r *http.Request) {
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

	var exports []struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	err = json.NewDecoder(r.Body).Decode(&exports)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var exportList []database.Export
	for _, export := range exports {
		exportList = append(exportList, database.Export{
			Name:     export.Name,
			Quantity: export.Quantity,
		})
	}

	err = services.EditExports(countryID, exportList)
	if err != nil {
		http.Error(w, "Error editing exports", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Exports edited successfully"))
}
