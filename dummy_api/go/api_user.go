package swagger

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = []User{
	{
		Id:          1,
		FirstName:   "John",
		Name:        "Doe",
		Password:    "password123",
		IdGraulande: "john123",
		MutuelleId:  1,
	},
	{
		Id:          2,
		FirstName:   "Jane",
		Name:        "Smith",
		Password:    "pass456",
		IdGraulande: "jane456",
		MutuelleId:  2,
	},
}

// UserAuthPost handles user authentication using idGraulande and password
func UserAuthPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Parse the request body to get user credentials
	var userCredentials User
	if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate user credentials (basic example, you might want to use a more secure authentication mechanism)
	for _, user := range users {
		if user.IdGraulande == userCredentials.IdGraulande && user.Password == userCredentials.Password {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// Authentication failed
	w.WriteHeader(http.StatusUnauthorized)
}

// UserGet retrieves the list of users
func UserGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// UserIdGet retrieves a user by ID
func UserIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the user ID from the request URL
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	if userIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert userIDStr to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the user with the given ID
	var result *User
	for _, user := range users {
		if user.Id == userID {
			result = &user
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

// UserIdDelete deletes a user by ID
func UserIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the user ID from the request URL
	vars := mux.Vars(r)
	userIDStr := vars["id"]
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

	// Find the index of the user with the given ID
	var index = -1
	for i, user := range users {
		if user.Id == userID {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Delete the user from the slice
	users = append(users[:index], users[index+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

// UserIdPut updates a user by ID
func UserIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the user ID from the request URL
	vars := mux.Vars(r)
	userIDStr := vars["id"]
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

	// Find the index of the user with the given ID
	var index = -1
	for i, user := range users {
		if user.Id == userID {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update the user with the new data
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Replace the existing user with the updated one
	users[index] = updatedUser

	err = json.NewEncoder(w).Encode(updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// UserPost creates a new user
func UserPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Parse the request body to create a new user
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Assign a new ID to the user
	newUser.Id = len(users) + 1

	// Add the new user to the slice
	users = append(users, newUser)

	err := json.NewEncoder(w).Encode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

// UserGraulandeIdGet retrieves a user by idGraulande
func UserGraulandeIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Extract the idGraulande from the request URL
	vars := mux.Vars(r)
	idGraulande := vars["id_graulande"]
	if idGraulande == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the user with the given idGraulande
	var result *User
	for _, user := range users {
		if user.IdGraulande == idGraulande {
			result = &user
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
