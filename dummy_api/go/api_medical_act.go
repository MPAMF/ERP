/*
 * Medical Records API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var medicalActs = []MedicalAct{
	{
		Id:                          1,
		UserId:                      1,
		HospitalId:                  1,
		MutuelleId:                  1,
		Metadata1:                   "test1",
		Metadata2:                   "test2",
		MontantTotal:                100,
		PriseEnChargeHopital:        2,
		PriseEnChargeMutuelle:       3,
		PriseEnChargePatient:        4,
		ConfirmationPaiementPatient: true,
		ConfirmationMutuelle:        false,
		ConfirmationRdv:             true,
		PourcentagePriseEnCharge:    2,
		Commentaire:                 "",
		DateCreation:                time.Now().AddDate(0, 0, -2),
		DatePrevue:                  time.Now().AddDate(0, 0, -1),
		DateVenue:                   time.Now().AddDate(0, 0, -1),
	}, {
		Id:                          2,
		UserId:                      1,
		HospitalId:                  1,
		MutuelleId:                  1,
		Metadata1:                   "test1",
		Metadata2:                   "test2",
		MontantTotal:                150,
		PriseEnChargeHopital:        2,
		PriseEnChargeMutuelle:       3,
		PriseEnChargePatient:        4,
		ConfirmationPaiementPatient: true,
		ConfirmationMutuelle:        false,
		ConfirmationRdv:             true,
		PourcentagePriseEnCharge:    2,
		Commentaire:                 "",
		DateCreation:                time.Now().AddDate(0, 0, -1),
		DatePrevue:                  time.Now().AddDate(0, 0, 10),
		DateVenue:                   time.Time{},
	},
}

func MedicalActGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(medicalActs)
}

func MedicalActIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID from the request URL
	// For example, assuming the URL is like "/medical-acts/{id}"
	// you can use a router library or parse it manually.
	// For simplicity, let's assume the ID is passed as a query parameter.
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

	// Find the medical act with the given ID
	var result *MedicalAct
	for _, act := range medicalActs {
		if act.Id == id {
			result = &act
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

// MedicalActIdDelete deletes a medical act by ID
func MedicalActIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID from the request URL
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

	// Find the index of the medical act with the given ID
	var index = -1
	for i, act := range medicalActs {
		if act.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Delete the medical act from the slice
	medicalActs = append(medicalActs[:index], medicalActs[index+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

// MedicalActIdPut updates a medical act by ID
func MedicalActIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID from the request URL
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

	// Find the index of the medical act with the given ID
	var index = -1
	for i, act := range medicalActs {
		if act.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update the medical act with the new data
	var updatedAct MedicalAct
	if err := json.NewDecoder(r.Body).Decode(&updatedAct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Replace the existing medical act with the updated one
	medicalActs[index] = updatedAct

	err = json.NewEncoder(w).Encode(updatedAct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// MedicalActUserIdGet retrieves medical acts for a specific user ID
func MedicalActUserIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the user ID from the request URL
	userIDStr := r.URL.Query().Get("userId")
	if userIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert userIDStr to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find all medical acts associated with the user ID
	var userMedicalActs []MedicalAct
	for _, act := range medicalActs {
		if act.UserId == userID {
			userMedicalActs = append(userMedicalActs, act)
		}
	}

	err = json.NewEncoder(w).Encode(userMedicalActs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// MedicalActPost creates a new medical act
func MedicalActPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Parse the request body to create a new medical act
	var newAct MedicalAct
	if err := json.NewDecoder(r.Body).Decode(&newAct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Assign a new ID to the medical act
	newAct.Id = len(medicalActs) + 1

	// Add the new medical act to the slice
	medicalActs = append(medicalActs, newAct)

	err := json.NewEncoder(w).Encode(newAct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}