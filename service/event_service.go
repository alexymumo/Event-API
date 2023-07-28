package service

import (
	"events/models"
	"events/repository"
)

func Save(event *models.Event) (*models.Event, error) {
	return repository.SaveEvent(event)
}
