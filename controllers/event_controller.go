package controllers

import (
	"encoding/json"
	"events/models"
	"events/service"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {

	}
	result, err := service.Save(&event)
	if err != nil {

	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(result)

}

func DeleteEventById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
