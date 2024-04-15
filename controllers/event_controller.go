package controllers

import (
	"encoding/json"
	"events/models"
	"events/responses"
	"events/service"
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
	fmt.Printf("Endpoint working")
}

type Error struct {
	Message string `json:"message"`
}

func (*controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	eventservice.Validate(&event)
	result, err := eventservice.Save(&event)
	if err != nil {
		panic(err.Error())
	}
	json, err := json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(201)
	fmt.Fprintf(w, "%s", json)
	responses.JSON(w, http.StatusCreated, result)

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
