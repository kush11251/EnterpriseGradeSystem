package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EnterpriseGradeSystem/pkg/models"
	"github.com/EnterpriseGradeSystem/pkg/services"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	fmt.Println("Controllers initialized")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	services := services.NewService()
	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}