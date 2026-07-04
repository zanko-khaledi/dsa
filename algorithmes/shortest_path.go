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

func Dijkestra(graph map[string][]Edge, start string) (map[string]int, map[string]string) {

	dist := make(map[string]int, len(graph))
	visited := make(map[string]bool)
	prev := make(map[string]string)

	for node := range graph {
		dist[node] = math.MaxInt
		prev[node] = ""
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
				prev[edge.To] = minNode
			}
		}
	}

	return dist, prev
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

func DijkestraWithHeap(graph map[string][]Edge, start string) (map[string]int, map[string]string) {

	dist := make(map[string]int)
	prev := make(map[string]string)

	for node := range graph {
		dist[node] = math.MaxInt
		prev[node] = ""
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

				prev[edge.To] = current.Node
			}
		}
	}

	return dist, prev
}

func BuildPath(edges map[string]string, start string, end string) []string {
	path := []string{}

	current := end

	for current != "" {

		path = append(path, current)

		if current == start {
			break
		}

		current = edges[current]
	}

	i, j := 0, len(path)-1

	for i <= j {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}

	return path
}
