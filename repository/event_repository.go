package repository

import (
	"database/sql"

	"events/models"
)

var db *sql.DB

func SaveEvent(event *models.Event) (*models.Event, error) {
	stmt, err := db.Prepare("INSERT INTO events(title,description,location) VALUES(?,?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := db.Exec(event.Title, event.Description, event.Location)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	event.EventID = int(id)
	return event, nil
}

func DeleteEvent(eventId int) (int64, error) {
	event, err := db.Exec("DELETE FROM events WHERE eventId=?", eventId)
	if err != nil {
		panic(err)
	}
	res, err := event.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, err
}

/*return slice of events*/

func FindAllEvents() ([]models.Event, error) {
	query, err := db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	var events []models.Event

	for query.Next() {
		var event models.Event
		err := query.Scan(
			&event.EventID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.CreatedAt,
			&event.UpdatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func UpdateEvent() {

}
