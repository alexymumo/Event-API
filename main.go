package main

import (
	"events/routes"

	_ "github.com/go-sql-driver/mysql"
)

/*
type Event struct {
	EventID     int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

var db *sql.DB
var err error

func createEvent(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO event(title, description, location) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	title := keyVal["title"]
	description := keyVal["description"]
	location := keyVal["location"]
	json.Unmarshal(body, &keyVal)
	_, err = stmt.Exec(title, description, location)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "post created")

}
*/

func main() {
	routes.Routes()
}
