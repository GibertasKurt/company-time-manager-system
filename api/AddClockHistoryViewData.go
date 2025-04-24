package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kurt/company-time-manager-system/models"
)

func AddClockHistoryAPIHandler(w http.ResponseWriter, r *http.Request) {
	var clockHistory models.ClockHistory
	err := json.NewDecoder(r.Body).Decode(&clockHistory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clockHistory.Save()

	uadminURL := "uadmin/admin/clockhistory/"
	jsonData, err := json.Marshal(clockHistory)
	if err != nil {
		http.Error(w, "Failed to marshal data for uadmin", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(uadminURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to send data to uadmin: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "uadmin responded with status: "+resp.Status, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(clockHistory)
}
