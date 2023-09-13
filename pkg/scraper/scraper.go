package scraper

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(link string) []string {

	url := "http://localhost:8080/wikipedia_en_all_nopic_2023-06/A/" + link
	res, err := http.Get(url)
	if err != nil {

		return []string{}
	}

	defer res.Body.Close()
	if res.StatusCode == 404 {
		return []string{}
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	content := doc.Find("#mw-content-text")

	var links []string
	content.Find("a").Each(func(i int, s *goquery.Selection) {
		href, found := s.Attr("href")

		if found && !strings.Contains(href, "http") && !strings.Contains(href, "#") && !strings.Contains(href, "geo:") && !strings.Contains(href, "news:") && !strings.Contains(href, "ftp:") {
			links = append(links, href)
		}
	})

	return links
}
