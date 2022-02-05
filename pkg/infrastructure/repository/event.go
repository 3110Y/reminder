package repository

import (
	"github.com/3110Y/reminder/pkg/Domain/models"
)

type Event struct {
	database database
	event    models.Event
}

func NewEvent(database database, event models.Event) *Event {
	return &Event{database: database, event: event}
}

func (e *Event) Create(event models.Event) (models.Event, error) {
	connection, err := e.database.Connection()
	if err != nil {
		return event, err
	}
	result := connection.Create(event)
	return event, result.Error
}

func (e *Event) Update(event interface{}, id uint) (models.Event, error) {
	connection, err := e.database.Connection()
	if err != nil {
		return e.event, err
	}
	e.event.ID = id
	result := connection.Model(&e.event).Updates(event)
	return e.event, result.Error
}

func (e *Event) Find(id uint) (models.Event, error) {
	connection, err := e.database.Connection()
	if err != nil {
		return e.event, err
	}
	result := connection.First(&e.event, id)
	return e.event, result.Error
}

func (e *Event) List(id uint, limit int) (map[int]models.Event, error) {
	mapEvent := make(map[int]models.Event)
	connection, err := e.database.Connection()
	if err != nil {
		return mapEvent, err
	}
	result := connection.Model(&e.event).Limit(limit).Where("id = ?", id).Take(mapEvent)
	return mapEvent, result.Error
}

func (e *Event) Delete(id uint) (err error) {
	connection, err := e.database.Connection()
	if err != nil {
		return err
	}
	result := connection.Delete(&models.Event{}, id)
	return result.Error
}
