package service

import (
	"github.com/3110Y/reminder/pkg/Domain/models"
	"github.com/3110Y/reminder/pkg/infrastructure/repository"
	"github.com/jinzhu/copier"
	"time"
)

type EventService struct {
	Event *repository.Event
}

type Event struct {
	FREQ       string    `json:"freq"`
	ByMonth    uint8     `json:"byMonth"`
	ByMonthDay uint8     `json:"byMonthDay"`
	ByDay      string    `json:"byDay"`
	BySetPos   int8      `json:"bySetPos"`
	Count      uint16    `json:"count"`
	Until      time.Time `json:"until"`
	Interval   uint16    `json:"interval"`
	Hook       string    `json:"hook"`
	StartDate  time.Time `json:"startDate"`
	NextDate   time.Time `json:"-"`
}

type EventWithID struct {
	Event
	ID uint
}

func NewEventService(event *repository.Event) *EventService {
	return &EventService{Event: event}
}

func (e *EventService) Add(event Event) (eventResponse EventWithID, err error) {
	model := models.Event{}
	err = copier.Copy(&model, event)
	if err != nil {
		return eventResponse, err
	}
	model, err = e.Event.Create(model)
	if err != nil {
		return eventResponse, err
	}
	err = copier.Copy(&eventResponse, model)
	return eventResponse, err
}

func (e *EventService) Edit(event Event, id uint) (eventResponse EventWithID, err error) {
	model, err := e.Event.Update(event, id)
	if err != nil {
		return eventResponse, err
	}
	err = copier.Copy(&eventResponse, model)
	return eventResponse, err
}

func (e *EventService) Delete(id uint) error {
	return e.Event.Delete(id)
}

func (e *EventService) Get(id uint) (eventResponse EventWithID, err error) {
	event, err := e.Event.Find(id)
	if err != nil {
		return eventResponse, err
	}
	err = copier.Copy(eventResponse, event)
	return eventResponse, err
}

func (e *EventService) List(id uint, limit int) (eventResponseList map[int]EventWithID, err error) {
	modelsEventList, err := e.Event.List(id, limit)
	if err != nil {
		return eventResponseList, err
	}
	err = copier.Copy(eventResponseList, modelsEventList)
	return eventResponseList, err
}
