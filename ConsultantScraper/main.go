package main

import (
	scraper "ConsultantScraper/src"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	webScraper := scraper.ConsultantScraper{}
	calendar := webScraper.GetSimpleCalendar(2023)

	serialized, err := json.Marshal(calendar)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	_ = ioutil.WriteFile("./Output/calendar.json", serialized, 0644)
}
