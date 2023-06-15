package src_webScraper

import (
	"time"
)

type BaseCalendar struct {
	Year int
}

func (s *BaseCalendar) GetMonthFromString(month string) time.Month {
	var values = map[string]time.Month{
		"январь":   time.January,
		"февраль":  time.February,
		"март":     time.March,
		"апрель":   time.April,
		"май":      time.May,
		"июнь":     time.June,
		"июль":     time.July,
		"август":   time.August,
		"сентябрь": time.September,
		"октябрь":  time.October,
		"ноябрь":   time.November,
		"декабрь":  time.December,
	}

	return values[month]
}

type SimpleCalendar struct {
	BaseCalendar
	Days []time.Time
}

type Calendar struct {
	BaseCalendar
	Days []Day
}

func (s *Calendar) CreateDate(dayType int, year int, month time.Month, day int) Day {
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	result := Day{
		Type: dayType,
		Date: date,
	}
	return result
}

type Day struct {
	Type int
	Date time.Time
}

const (
	Weekend int = iota
	Holiday
)
