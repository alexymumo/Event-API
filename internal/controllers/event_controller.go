package controllers

import (
	"events/internal/models"
	"events/internal/service"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	eventservice service.EventService
)

type EventController interface {
	Test(c *gin.Context)
	GetAllEvents(c *gin.Context)
	CreateEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
	UpdateEvent(c *gin.Context)
}

type controller struct{}

func EventControllerImpl(service service.EventService) EventController {
	eventservice = service
	return &controller{}
}

func (*controller) Test(c *gin.Context) {
	fmt.Println("End Points Working")
}

type Error struct {
	Message string `json:"message"`
}

func (*controller) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err := eventservice.Validate(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := eventservice.Save(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*controller) GetAllEvents(c *gin.Context) {
	/*
		w.Header().Set("Content-Type", "application/json")
		events, err := eventservice.FindAllEvents()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`Error:"Error occurred"`))
		}
		_ = json.NewEncoder(w).Encode(events)
	*/
}

func (*controller) DeleteEvent(c *gin.Context) {
	//w.Header().Set("Content-Type", "application/json")

}

func (*controller) UpdateEvent(c *gin.Context) {
	//w.Header().Set("Content-Type", "application/json")

}
