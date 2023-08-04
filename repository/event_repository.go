package repository

import (
	"events/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

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
	return 0, nil
	/*
		event, err := db.Exec(`DELETE FROM events WHERE eventId=?`, eventId)
		if err != nil {
			panic(err)
		}
		res, err := event.RowsAffected()
		if err != nil {
			return 0, err
		}
		return res, err
	*/
}

/*return slice of events*/

func (*repository) FindAllEvents() ([]models.Event, error) {
	var events []models.Event
	return events, nil

	/*query, err := db.Query(`SELECT * FROM events`)
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
	*/
}

func (*repository) Update(event *models.Event, id int64) (*models.Event, error) {
	return event, nil
	/*stmt, err := db.Prepare(`UPDATE events SET title=?,description=?, location=? WHERE id=?;`)
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
	*/
}

func (*repository) Save(event *models.Event) (*models.Event, error) {
	err := db.Debug().Model(&models.Event{}).Create(&event).Error
	if err != err {
		return &models.Event{}, err
	}
	return event, nil
}
