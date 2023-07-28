package routes

import (
	"events/controllers"
)

var EventsRoutes = func() {
	router.HandleFunc("v1/event", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("v1/event/{id}", controllers.DeleteEventById).Methods("DELETE")
	router.HandleFunc("v1/events", controllers.GetEvents).Methods("GET")
	router.HandleFunc("v1/", controllers.UpdateEvent).Methods("PUT")

	router.HandleFunc()
}
