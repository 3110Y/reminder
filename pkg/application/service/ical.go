package service

import "errors"

type ICal struct {
	event *Event
}

func (c *ICal) fill(event *Event) (err error) {
	c.event = event
	switch event.FREQ {
	case "YEARLY":
		return c.yearly()
	case "MONTHLY":
		return c.monthly()
	case "WEEKLY":
		return c.weekly()
	case "DAILY":
		return c.daily()
	default:
		err = errors.New("не известный тип FREQ")
	}
	return
}

func (c *ICal) yearly() (err error) {
	return
}

func (c *ICal) monthly() (err error) {
	return
}

func (c *ICal) weekly() (err error) {
	return
}

func (c *ICal) daily() (err error) {
	return
}
