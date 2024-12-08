package handlers

import (
	"encoding/json"
	"go-sqlite-backend/database"
	"go-sqlite-backend/services"
	"net/http"
	"strconv"
)

func EditImportsHandler(w http.ResponseWriter, r *http.Request) {
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

	var imports []struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	err = json.NewDecoder(r.Body).Decode(&imports)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var importList []database.Import
	for _, importItem := range imports {
		importList = append(importList, database.Import{
			Name:     importItem.Name,
			Quantity: importItem.Quantity,
		})
	}

	err = services.EditImports(countryID, importList)
	if err != nil {
		http.Error(w, "Error editing imports", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Imports edited successfully"))
}
