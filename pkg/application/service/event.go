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

func NewEventService(event *repository.Event) *EventService {
	return &EventService{Event: event}
}

func (e EventService) Add(event *Event) (err error) {
	model := &models.Event{}
	err = copier.Copy(model, event)
	if err != nil {
		return err
	}
	return e.Event.Create(*model)
}
