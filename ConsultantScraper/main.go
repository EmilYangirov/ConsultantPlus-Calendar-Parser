package main

import (
	scraper "ConsultantScraper/src"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	webScraper := scraper.ConsultantScraper{}
	year := 2024
	calendar := webScraper.GetSimpleCalendar(year)

	serialized, err := json.Marshal(calendar)

	name := "calendar_" + strconv.Itoa(year) + ".json"

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	_ = ioutil.WriteFile("./Output/"+name, serialized, 0644)
}
