package main

import (
	. "WikiShortestPath/pkg/article"
	. "WikiShortestPath/pkg/bfs"
	"WikiShortestPath/pkg/scraper"
	"fmt"
	"os"
	"sync"
)

func main() {

	// Start Article
	start := "Pineapple"

	// Target Article
	target := "Vin_Diesel"

	bfs := new(BFS)
	bfs.Visited = make(map[string]int)

	links := scraper.Scraper(start)
	first := Article{Title: start, History: []string{}, Links: links}

	bfs.Enqueue(first)

	for {

		ch := make(chan Article)
		var wg sync.WaitGroup

		next := bfs.Dequeue()

		for _, link := range next.Links {

			_, in_visited_map := bfs.Visited[link]
			if in_visited_map {
				continue
			}

			// Takes the history of the parents, and adds the current title. This becomes the path
			history := append(next.History, next.Title)

			wg.Add(1)
			go get_links_for_article(link, history, ch, &wg)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for result := range ch {
			if result.Title == target {
				finish_and_exit(result.History, result.Title)
			}

			bfs.AddToVisited(result.Title)
			bfs.Enqueue(result)
		}

	}

}

func get_links_for_article(l string, h []string, ch chan<- Article, wg *sync.WaitGroup) {

	defer wg.Done()
	links := scraper.Scraper(l)
	article := Article{Title: l, History: h, Links: links}
	ch <- article
}

func finish_and_exit(h []string, l string) {
	history := append(h, l)
	fmt.Println("Found it, the path is: ")
	for i, h := range history {
		if i == 0 {
			fmt.Print("", h)
			continue
		}
		fmt.Print(" -> ", h)
	}
	fmt.Println("")
	os.Exit(0)
}
