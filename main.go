package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	res, err := http.Get("http://localhost:8080/wikipedia_en_all_nopic_2023-06/A/Pineapple")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#mw-content-text").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			href, found := s.Attr("href")
			if found && !strings.Contains(href, "http") && !strings.Contains(href, "#cite") {
				fmt.Printf("Review %d: %s\n\n", i, href)
			}
		})
	})
}
