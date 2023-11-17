package swagger

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var medicalActResults = []MedicalActResult{
	{
		Id:             1,
		MedicalActId:   1,
		FileName:       "file1.txt",
		FileData:       "VGhpcyBpcyBhIHRlc3QgZmlsZQo=",
		FileBinaryData: []byte("This is a test file\n"),
	},
}

// MedicalActIdResultsGet retrieves the list of medical act results for a specific medical act
func MedicalActIdResultsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID from the request URL
	vars := mux.Vars(r)
	medicalActIDStr := vars["id"]
	if medicalActIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert medicalActIDStr to int
	medicalActID, err := strconv.Atoi(medicalActIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Filter medical act results by the given medical act ID
	var resultsForMedicalAct []MedicalActResult
	for _, result := range medicalActResults {
		if result.MedicalActId == medicalActID {
			resultsForMedicalAct = append(resultsForMedicalAct, result)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultsForMedicalAct)
}

// MedicalActIdResultsPost creates a new medical act result for a specific medical act
func MedicalActIdResultsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID from the request URL
	vars := mux.Vars(r)
	medicalActIDStr := vars["id"]
	if medicalActIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert medicalActIDStr to int
	medicalActID, err := strconv.Atoi(medicalActIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the request body to create a new medical act result
	var newMedicalActResult MedicalActResult
	if err := json.NewDecoder(r.Body).Decode(&newMedicalActResult); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Assign a new ID to the medical act result
	newMedicalActResult.Id = len(medicalActResults) + 1
	newMedicalActResult.MedicalActId = medicalActID

	// Add the new medical act result to the slice
	medicalActResults = append(medicalActResults, newMedicalActResult)

	err = json.NewEncoder(w).Encode(newMedicalActResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

// MedicalActIdResultsResultIdGet retrieves a specific medical act result by ID for a given medical act
func MedicalActIdResultsResultIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID and result ID from the request URL
	vars := mux.Vars(r)
	medicalActIDStr := vars["id"]
	resultIDStr := vars["result_id"]

	if medicalActIDStr == "" || resultIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert medicalActIDStr and resultIDStr to int
	medicalActID, err := strconv.Atoi(medicalActIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resultID, err := strconv.Atoi(resultIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the medical act result with the given IDs
	var result *MedicalActResult
	for _, medicalActResult := range medicalActResults {
		if medicalActResult.MedicalActId == medicalActID && medicalActResult.Id == resultID {
			result = &medicalActResult
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

// MedicalActIdResultsResultIdPut updates a specific medical act result by ID for a given medical act
func MedicalActIdResultsResultIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID and result ID from the request URL
	vars := mux.Vars(r)
	medicalActIDStr := vars["id"]
	resultIDStr := vars["result_id"]

	if medicalActIDStr == "" || resultIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert medicalActIDStr and resultIDStr to int
	medicalActID, err := strconv.Atoi(medicalActIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resultID, err := strconv.Atoi(resultIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the index of the medical act result with the given IDs
	var index = -1
	for i, medicalActResult := range medicalActResults {
		if medicalActResult.MedicalActId == medicalActID && medicalActResult.Id == resultID {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update the medical act result with the new data
	var updatedMedicalActResult MedicalActResult
	if err := json.NewDecoder(r.Body).Decode(&updatedMedicalActResult); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Replace the existing medical act result with the updated one
	medicalActResults[index] = updatedMedicalActResult

	err = json.NewEncoder(w).Encode(updatedMedicalActResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// MedicalActIdResultsResultIdDelete deletes a specific medical act result by ID for a given medical act
func MedicalActIdResultsResultIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the medical act ID and result ID from the request URL
	vars := mux.Vars(r)
	medicalActIDStr := vars["id"]
	resultIDStr := vars["result_id"]

	if medicalActIDStr == "" || resultIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert medicalActIDStr and resultIDStr to int
	medicalActID, err := strconv.Atoi(medicalActIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resultID, err := strconv.Atoi(resultIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the index of the medical act result with the given IDs
	var index = -1
	for i, medicalActResult := range medicalActResults {
		if medicalActResult.MedicalActId == medicalActID && medicalActResult.Id == resultID {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Delete the medical act result from the slice
	medicalActResults = append(medicalActResults[:index], medicalActResults[index+1:]...)

	w.WriteHeader(http.StatusNoContent)
}
