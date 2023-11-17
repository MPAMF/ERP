package swagger

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var mutuelles = []Mutuelle{
	{
		Id:   1,
		Name: "Mutuelle A",
	},
	{
		Id:   2,
		Name: "Mutuelle B",
	},
}

// MutuelleGet retrieves the list of mutuelles
func MutuelleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mutuelles)
}

// MutuelleIdGet retrieves a mutuelle by ID
func MutuelleIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the mutuelle ID from the request URL
	vars := mux.Vars(r)
	mutuelleIDStr := vars["id"]
	if mutuelleIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert mutuelleIDStr to int
	mutuelleID, err := strconv.Atoi(mutuelleIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the mutuelle with the given ID
	var result *Mutuelle
	for _, mutuelle := range mutuelles {
		if mutuelle.Id == mutuelleID {
			result = &mutuelle
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

// MutuelleIdDelete deletes a mutuelle by ID
func MutuelleIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the mutuelle ID from the request URL
	vars := mux.Vars(r)
	mutuelleIDStr := vars["id"]
	if mutuelleIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert mutuelleIDStr to int
	mutuelleID, err := strconv.Atoi(mutuelleIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the index of the mutuelle with the given ID
	var index = -1
	for i, mutuelle := range mutuelles {
		if mutuelle.Id == mutuelleID {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Delete the mutuelle from the slice
	mutuelles = append(mutuelles[:index], mutuelles[index+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

// MutuelleIdPut updates a mutuelle by ID
func MutuelleIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the mutuelle ID from the request URL
	vars := mux.Vars(r)
	mutuelleIDStr := vars["id"]
	if mutuelleIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert mutuelleIDStr to int
	mutuelleID, err := strconv.Atoi(mutuelleIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the index of the mutuelle with the given ID
	var index = -1
	for i, mutuelle := range mutuelles {
		if mutuelle.Id == mutuelleID {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update the mutuelle with the new data
	var updatedMutuelle Mutuelle
	if err := json.NewDecoder(r.Body).Decode(&updatedMutuelle); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Replace the existing mutuelle with the updated one
	mutuelles[index] = updatedMutuelle

	err = json.NewEncoder(w).Encode(updatedMutuelle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// MutuellePost creates a new mutuelle
func MutuellePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Parse the request body to create a new mutuelle
	var newMutuelle Mutuelle
	if err := json.NewDecoder(r.Body).Decode(&newMutuelle); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Assign a new ID to the mutuelle
	newMutuelle.Id = len(mutuelles) + 1

	// Add the new mutuelle to the slice
	mutuelles = append(mutuelles, newMutuelle)

	err := json.NewEncoder(w).Encode(newMutuelle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
