package src_webScraper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	goquery "github.com/PuerkitoBio/goquery"
)

type ConsultantScraper struct {
	scraper
}

func (s *ConsultantScraper) createUrl(year int) string {
	baseUrl := "https://www.consultant.ru/law/ref/calendar/proizvodstvennye/"

	return baseUrl + strconv.Itoa(year) + "/"
}

func (s *ConsultantScraper) GetCalendarWithTypes(year int) Calendar {
	url := s.createUrl(year)
	doc := s.loadPage(url)

	var calendar Calendar
	calendar.Year = year

	doc.Find(".cal").Each(func(i int, s *goquery.Selection) {
		stringedMonth := s.Find(".month").Text()

		month := calendar.GetMonthFromString(strings.ToLower(stringedMonth))
		s.Find("td[class=weekend]").Each(func(j int, sel *goquery.Selection) {
			day, err := strconv.Atoi(sel.Text())

			if err == nil {
				calendar.Days = append(calendar.Days, calendar.CreateDate(Weekend, year, month, day))
			}
		})

		s.Find(".holiday.weekend").Each(func(j int, sel *goquery.Selection) {
			day, err := strconv.Atoi(sel.Text())
			if err != nil {
				fmt.Println(sel.Text())
			}

			if err == nil {
				calendar.Days = append(calendar.Days, calendar.CreateDate(Holiday, year, month, day))
			}
		})
	})
	fmt.Println(len(calendar.Days))
	return calendar
}

func (s *ConsultantScraper) GetSimpleCalendar(year int) SimpleCalendar {
	url := s.createUrl(year)
	doc := s.loadPage(url)

	var calendar SimpleCalendar
	calendar.Year = year

	doc.Find(".cal").Each(func(i int, s *goquery.Selection) {
		stringedMonth := s.Find(".month").Text()

		month := calendar.GetMonthFromString(strings.ToLower(stringedMonth))
		s.Find(".weekend").Each(func(j int, sel *goquery.Selection) {
			day, err := strconv.Atoi(sel.Text())

			if err == nil {
				day := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
				calendar.Days = append(calendar.Days, day)
			}
		})
	})

	fmt.Println(len(calendar.Days))
	return calendar
}
