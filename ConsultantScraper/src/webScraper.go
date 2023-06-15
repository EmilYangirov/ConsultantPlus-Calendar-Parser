package src_webScraper

import (
	"fmt"
	"net/http"

	goquery "github.com/PuerkitoBio/goquery"
)

type iScraper interface {
	loadPage(url string) *goquery.Document
	createUrl(keyword []string) string
}

type scraper struct {
}

func (base *scraper) loadPage(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 && res.StatusCode != 300 {
		fmt.Println(res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return doc
}
