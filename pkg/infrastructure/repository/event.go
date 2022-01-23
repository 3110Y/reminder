package repository

import (
	"github.com/3110Y/reminder/pkg/Domain/models"
)

type Event struct {
	database database
}

func NewEvent(database database) *Event {
	return &Event{database: database}
}

func (e Event) Create(event models.Event) (err error) {
	connection, err := e.database.Connection()
	if err != nil {
		return err
	}
	result := connection.Create(event)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (e Event) Update(event interface{}) (err error) {
	connection, err := e.database.Connection()
	if err != nil {
		return err
	}
	result := connection.Save(event)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (e Event) Find(event *models.Event, id int) (err error) {
	connection, err := e.database.Connection()
	if err != nil {
		return err
	}
	result := connection.First(event, id)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (e Event) list(model interface{}, id, limit int) (map[string]models.Event, error) {
	mapEvent := make(map[string]models.Event)
	connection, err := e.database.Connection()
	if err != nil {
		return mapEvent, err
	}
	result := connection.Model(model).Limit(limit).Where("id = ?", id).Take(mapEvent)
	if result.Error != nil {
		return mapEvent, err
	}
	return mapEvent, nil
}

func (e Event) delete(id int) (err error) {
	connection, err := e.database.Connection()
	if err != nil {
		return err
	}
	result := connection.Delete(&models.Event{}, id)
	if result.Error != nil {
		return result.Error
	}
	return
}
