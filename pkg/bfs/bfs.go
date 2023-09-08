package bfs

import . "WikiShortestPath/pkg/article"

type BFS struct {
	Queue   []Article
	Visited []string
}

// Enqueue puts an item at the end of the queue
func (b *BFS) Enqueue(a Article) {
	b.Queue = append(b.Queue, a)
}

// Dequeue takes an item from the front of the queue
func (b *BFS) Dequeue() Article {
	first := b.Queue[0]
	b.Queue = b.Queue[1:]

	return first
}

func (b *BFS) AddToVisited(s string) {
	b.Visited = append(b.Visited, s)
}
