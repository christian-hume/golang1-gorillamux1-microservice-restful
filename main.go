package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/golang/api/v1/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/golang/api/v1/users", CreateUser).Methods("POST")
	router.HandleFunc("/golang/api/v1/users/{id}", GetUserByID).Methods("GET")
	router.HandleFunc("/golang/api/v1/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/golang/api/v1/users/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
