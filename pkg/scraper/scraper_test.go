package scraper

import "testing"

func TestScraper(t *testing.T) {

	links := Scraper("Pineapple")

	got := len(links)
	want := 442

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}
