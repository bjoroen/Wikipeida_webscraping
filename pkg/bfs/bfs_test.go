package bfs

import (
	. "WikiShortestPath/pkg/article"
	"testing"
)

func TestBfsEnqueueDequeue(t *testing.T) {
	bfs := new(BFS)

	a := Article{Title: "first", History: []string{"one", "two"}, Links: []string{"three", "four"}}
	b := Article{Title: "second", History: []string{"1", "2"}, Links: []string{"3", "4"}}

	bfs.Enqueue(a)
	bfs.Enqueue(b)

	got := bfs.Dequeue()
	want := a

	if got.Title != want.Title {
		t.Errorf("got %s, wanted %s", got, want)
	}

	for i, e := range got.History {
		if e != want.History[i] {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}

	for i, e := range got.Links {
		if e != want.Links[i] {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}

}

func TestDequeueRemoves(t *testing.T) {
	bfs := new(BFS)

	a := Article{Title: "first", History: []string{"one", "two"}, Links: []string{"three", "four"}}
	b := Article{Title: "second", History: []string{"1", "2"}, Links: []string{"3", "4"}}

	bfs.Enqueue(a)
	bfs.Enqueue(b)

	bfs.Dequeue()

	got := len(bfs.Queue)
	want := 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func TestAddToVisited(t *testing.T) {
	bfs := new(BFS)
	bfs.Visited = make(map[string]int)

	bfs.AddToVisited("helloWorld")

	got := bfs.Visited["helloWorld"]
	want := 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
