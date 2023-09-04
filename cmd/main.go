package main

import (
	"WikiShortestPath/pkg/scraper"
	"fmt"
)

func main() {
	s := scraper.Scraper("Pineapple")
	fmt.Println(s)
}
