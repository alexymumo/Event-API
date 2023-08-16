package controllers

import (
	"encoding/json"
	"events/models"
	"events/service"
	"events/util"
	"fmt"
	"net/http"
)

var (
	eventservice service.EventService
)

type EventController interface {
	Test(w http.ResponseWriter, r *http.Request)
	GetAllEvents(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func EventControllerImpl(service service.EventService) EventController {
	eventservice = service
	return &controller{}
}

func (*controller) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("End Points Working")
}

type Error struct {
	Message string `json:"message"`
}

func (*controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.ServiceError{Message: "Wrong data format"})
	}

	valErr := eventservice.Validate(&event)
	if valErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.ServiceError{Message: valErr.Error()})
	}
	result, err1 := eventservice.Save(&event)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.ServiceError{Message: "Error encountered while saving..."})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*controller) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	events, err := eventservice.FindAllEvents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Error:"Error occurred"`))
	}
	_ = json.NewEncoder(w).Encode(events)
}

func (*controller) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (*controller) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
