package bfs

import . "WikiShortestPath/pkg/article"

type BFS struct {
	queue   []Article
	visited []string
}

func (b *BFS) Enqueue(a Article) {
	b.queue = append(b.queue, a)
}

func (b *BFS) Dequeue() Article {
	first := b.queue[0]
	b.queue = b.queue[:1]

	return first
}

func (b *BFS) AddToVisited(s string) {
	b.visited = append(b.visited, s)
}
