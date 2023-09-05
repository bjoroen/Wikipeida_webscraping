package bfs

import (
	. "WikiShortestPath/pkg/article"
	"testing"
)

func TestBfsEnqueueDequeue(t *testing.T) {
	bfs := new(BFS)

	a := Article{Current: "first", History: []string{"one", "two"}, Links: []string{"three", "four"}}
	b := Article{Current: "second", History: []string{"1", "2"}, Links: []string{"3", "4"}}

	bfs.Enqueue(a)
	bfs.Enqueue(b)

	got := bfs.Dequeue()
	want := a

	if got.Current != want.Current {
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

func TestAddToVisited(t *testing.T) {
	bfs := new(BFS)

	bfs.AddToVisited("helloWorld")

	got := bfs.visited[0]
	want := "helloWorld"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
