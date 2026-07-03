package algorithmes

import (
	"container/heap"
	"math"
)

type Edge struct {
	To     string
	Weight int
}

type Item struct {
	Node string
	Dist int
}

type PriorityQueue []Item

func Dijkestra(graph map[string][]Edge, start string) map[string]int {

	dist := make(map[string]int, len(graph))
	visited := make(map[string]bool)

	for node := range graph {
		dist[node] = math.MaxInt
	}

	dist[start] = 0

	for range graph {

		minNode := ""
		minDist := math.MaxInt

		for node, d := range dist {
			if !visited[node] && d < minDist {
				minDist = d
				minNode = node
			}
		}

		if minNode == "" {
			break
		}

		visited[minNode] = true

		for _, edge := range graph[minNode] {

			newDist := dist[minNode] + edge.Weight

			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
			}
		}
	}

	return dist
}

/* Heap methods intialazing */

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	*pq = old[:n-1]

	return item
}

func DijkestraWithHeap(graph map[string][]Edge, start string) map[string]int {

	dist := make(map[string]int, len(graph))

	for node := range graph {
		dist[node] = math.MaxInt
	}

	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{
		Node: start,
		Dist: 0,
	})

	for pq.Len() > 0 {

		current := heap.Pop(pq).(Item)

		if current.Dist > dist[current.Node] {
			continue
		}

		for _, edge := range graph[current.Node] {

			newDist := current.Dist + edge.Weight

			if newDist < dist[edge.To] {
				dist[edge.To] = newDist

				heap.Push(pq, Item{
					Node: edge.To,
					Dist: newDist,
				})
			}
		}
	}

	return dist
}
