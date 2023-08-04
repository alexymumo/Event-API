package main

import (
	"events/controllers"
	"events/models"
	"events/repository"
	"events/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	repo          repository.EventRepository  = repository.EventRepsotoryImpl()
	eventservices service.EventService        = service.EventServiceImpl(repo)
	controller    controllers.EventController = controllers.EventControllerImpl(eventservices)
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/event", controller.CreateEvent).Methods("POST")

	test := "root:@tcp(127.0.0.1:3306)/events?charset=utf8&parseTime=True&loc=Local"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("mysql", test)

	if err != nil {
		fmt.Printf("Cannot connect to database")
	} else {
		fmt.Printf("Connected successfully to database")
	}
	db.AutoMigrate(&models.Event{}, &models.User{})

	http.ListenAndServe(":8000", nil)
}
