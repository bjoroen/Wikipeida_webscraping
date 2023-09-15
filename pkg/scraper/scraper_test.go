package scraper

import (
	"testing"
)

func TestScraper(t *testing.T) {

	links := Scraper("Pineapple")

	got := len(links)
	want := 435

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func BenchmarkScraper(b *testing.B) {
	Scraper("Pineapple")
}
