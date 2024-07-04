package repository

import (
	"database/sql"
	"events/internal/models"
	"events/pkg/database"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	database.Connect()
	db = database.GetDb()
}

type EventRepository interface {
	Save(event *models.Event) (*models.Event, error)
	Delete(id int64) (int64, error)
	Update(event *models.Event, id int64) (*models.Event, error)
	FindAllEvents() ([]models.Event, error)
}

type repository struct{}

func EventRepsotoryImpl() EventRepository {
	return &repository{}
}

func (*repository) Delete(eventId int64) (int64, error) {
	event, err := db.Exec(`DELETE FROM events WHERE eventId=?`, eventId)
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

func (*repository) FindAllEvents() ([]models.Event, error) {
	query, err := db.Query(`SELECT * FROM events`)
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
			&event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}

func (*repository) Update(event *models.Event, id int64) (*models.Event, error) {
	stmt, err := db.Prepare(`UPDATE events SET title=?,description=?, location=? WHERE id=?;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = db.Exec(event.Title, event.Description, event.Location, id)
	if err != nil {
		return nil, err
	}
	event.EventID = int(id)
	return event, err

}

func (*repository) Save(event *models.Event) (*models.Event, error) {
	stmt, err := db.Prepare("INSERT INTO event(title, description, location) VALUES(?,?,?)")
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
