package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)


func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var reqBody previewTripRequest
	if err := json.Decoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" || reqBody.Pickup.Latitude == 0 || reqBody.Pickup.Longitude == 0 || reqBody.Destination.Latitude == 0 || reqBody.Destination.Longitude == 0 {
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}
	// Call the trip service to get trip preview
	response := contracts.APIResponse{Data: "ok"}
	writeJSON(w, http.StatusOK, response)
}