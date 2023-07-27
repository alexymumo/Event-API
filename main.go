package main

import (
	"events/db"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	router := mux.NewRouter()
	http.ListenAndServe("8080", router)
}
