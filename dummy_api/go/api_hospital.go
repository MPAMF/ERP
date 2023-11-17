package swagger

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var hospitals = []Hospital{
	{
		Id:   1,
		Name: "Hospital A",
	},
	{
		Id:   2,
		Name: "Hospital B",
	},
}

// HospitalGet retrieves the list of hospitals
func HospitalGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hospitals)
}

// HospitalIdGet retrieves a hospital by ID
func HospitalIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the hospital ID from the request URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert idStr to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the hospital with the given ID
	var result *Hospital
	for _, hospital := range hospitals {
		if hospital.Id == id {
			result = &hospital
			break
		}
	}

	if result == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// HospitalIdDelete deletes a hospital by ID
func HospitalIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the hospital ID from the request URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert idStr to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the index of the hospital with the given ID
	var index = -1
	for i, hospital := range hospitals {
		if hospital.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Delete the hospital from the slice
	hospitals = append(hospitals[:index], hospitals[index+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

// HospitalIdPut updates a hospital by ID
func HospitalIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the hospital ID from the request URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert idStr to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the index of the hospital with the given ID
	var index = -1
	for i, hospital := range hospitals {
		if hospital.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update the hospital with the new data
	var updatedHospital Hospital
	if err := json.NewDecoder(r.Body).Decode(&updatedHospital); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Replace the existing hospital with the updated one
	hospitals[index] = updatedHospital

	err = json.NewEncoder(w).Encode(updatedHospital)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// HospitalPost creates a new hospital
func HospitalPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Parse the request body to create a new hospital
	var newHospital Hospital
	if err := json.NewDecoder(r.Body).Decode(&newHospital); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Assign a new ID to the hospital
	newHospital.Id = len(hospitals) + 1

	// Add the new hospital to the slice
	hospitals = append(hospitals, newHospital)

	err := json.NewEncoder(w).Encode(newHospital)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
