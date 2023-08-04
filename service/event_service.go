package service

import (
	"errors"
	"events/models"
	"events/repository"
)

var (
	repo repository.EventRepository
)

type service struct{}

type EventService interface {
	Validate(event *models.Event) error
	Save(event *models.Event) (*models.Event, error)
	Delete(id int64) (int64, error)
	Update(event *models.Event, id int64) (*models.Event, error)
	FindAllEvents() ([]models.Event, error)
}

func EventServiceImpl(repository repository.EventRepository) EventService {
	repo = repository
	return &service{}

}

func (*service) Validate(event *models.Event) error {
	if event == nil {
		err := errors.New("cannot be empty")
		return err
	}

	if event.Title == "" {
		err := errors.New("cannot be empty")
		return err
	}
	if event.Description == "" {
		err := errors.New("cannot be empty")
		return err
	}
	if event.Location == "" {
		err := errors.New("cannot be empty")
		return err
	}
	return nil
}

func (*service) Save(event *models.Event) (*models.Event, error) {
	return repo.Save(event)
}

func (*service) Delete(id int64) (int64, error) {
	return repo.Delete(id)
}

func (*service) Update(event *models.Event, id int64) (*models.Event, error) {
	return repo.Update(event, id)

}

func (*service) FindAllEvents() ([]models.Event, error) {
	return repo.FindAllEvents()

}
