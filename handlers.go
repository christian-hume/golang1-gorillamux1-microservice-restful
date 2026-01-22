package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User represents a user model in memory
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Response standard structure for all API responses
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var users []User
var nextID = 1

// GetAllUsers - GET /users
func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	json.NewEncoder(response).Encode(Response{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

// CreateUser - POST /golang/api/v1/users
func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Response{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	user.ID = nextID
	nextID++
	users = append(users, user)

	json.NewEncoder(response).Encode(Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    user,
	})
}

// GetUserByID - GET /users/{id}
func GetUserByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Response{
			Status:  "error",
			Message: "Invalid user ID",
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(response).Encode(Response{
				Status:  "success",
				Message: "User found",
				Data:    user,
			})
			return
		}
	}

	response.WriteHeader(http.StatusNotFound)
	json.NewEncoder(response).Encode(Response{
		Status:  "error",
		Message: "User not found",
	})
}

// UpdateUser - PUT /users/{id}
func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Response{
			Status:  "error",
			Message: "Invalid user ID",
		})
		return
	}

	for index, user := range users {
		if user.ID == id {
			err := json.NewDecoder(request.Body).Decode(&users[index])
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(response).Encode(Response{
					Status:  "error",
					Message: "Invalid request body",
				})
				return
			}
			users[index].ID = id

			json.NewEncoder(response).Encode(Response{
				Status:  "success",
				Message: "User updated successfully",
				Data:    users[index],
			})
			return
		}
	}

	response.WriteHeader(http.StatusNotFound)
	json.NewEncoder(response).Encode(Response{
		Status:  "error",
		Message: "User not found",
	})
}

// DeleteUser - DELETE /users/{id}
func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Response{
			Status:  "error",
			Message: "Invalid user ID",
		})
		return
	}

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(response).Encode(Response{
				Status:  "success",
				Message: "User deleted successfully",
			})
			return
		}
	}

	response.WriteHeader(http.StatusNotFound)
	json.NewEncoder(response).Encode(Response{
		Status:  "error",
		Message: "User not found",
	})
}
