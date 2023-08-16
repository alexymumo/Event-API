package routes

import (
	"events/controllers"
	"events/repository"
	"events/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	repo          repository.EventRepository  = repository.EventRepsotoryImpl()
	eventservices service.EventService        = service.EventServiceImpl(repo)
	controller    controllers.EventController = controllers.EventControllerImpl(eventservices)
)

var Routes = func() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env variables")
	}
	router := mux.NewRouter()
	router.HandleFunc("/test", controller.Test)
	router.HandleFunc("/event", controller.CreateEvent)
	router.HandleFunc("/events", controller.GetAllEvents)
	http.ListenAndServe(":8080", router)
}
