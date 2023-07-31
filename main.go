package main

import (
	"events/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/event", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/event/{id}", controllers.DeleteEventById).Methods("DELETE")
	router.HandleFunc("/events", controllers.GetEvents).Methods("GET")
	router.HandleFunc("/event", controllers.UpdateEvent).Methods("PUT")
	http.ListenAndServe(":8080", router)
}
