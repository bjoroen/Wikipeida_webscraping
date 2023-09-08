package main

import (
	. "WikiShortestPath/pkg/article"
	. "WikiShortestPath/pkg/bfs"
	"WikiShortestPath/pkg/scraper"
	"fmt"
	"os"
	"slices"
)

func main() {

	start := "Zambia"
	target := "Old_French"

	bfs := new(BFS)

	links := scraper.Scraper(start)
	first := Article{Title: start, History: []string{}, Links: links}

	bfs.Enqueue(first)

	for {

		next := bfs.Dequeue()

		for _, link := range next.Links {

			history := append(next.History, next.Title)

			if link == target {
				history := append(history, link)
				fmt.Println("Found it, the path is: ")
				for i, h := range history {
					if i == 0 {
						fmt.Print("", h)
					}
					fmt.Print(" -> ", h)
				}
				os.Exit(0)
			}

			if slices.Contains(bfs.Visited, link) {
				continue
			}

			links := scraper.Scraper(link)
			article := Article{Title: link, History: history, Links: links}

			bfs.Enqueue(article)
			bfs.AddToVisited(link)

		}

	}

}
