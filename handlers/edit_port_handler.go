package handlers

import (
	"encoding/json"
	"go-sqlite-backend/database"
	"go-sqlite-backend/services"
	"net/http"
	"strconv"
)

func EditPortsHandler(w http.ResponseWriter, r *http.Request) {
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

	var ports []struct {
		Name string `json:"name"`
	}

	err = json.NewDecoder(r.Body).Decode(&ports)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var portList []database.Port
	for _, port := range ports {
		portList = append(portList, database.Port{
			Name: port.Name,
		})
	}

	err = services.EditPorts(countryID, portList)
	if err != nil {
		http.Error(w, "Error editing ports", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ports edited successfully"))
}
