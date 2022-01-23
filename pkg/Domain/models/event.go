package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	FREQ       string    //тип
	ByMonth    uint8     //месяц
	ByMonthDay uint8     //день
	ByDay      string    //воскресенье, enum
	BySetPos   int8      // какой по очереди, -1 последний
	Count      uint16    //кол-во повторений
	Until      time.Time // Дата окончания
	Interval   uint16    // итервал пракращения
	NextDate   time.Time
	Hook       string
	StartDate  time.Time
}
