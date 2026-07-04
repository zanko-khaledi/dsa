package main

import (
	"fmt"
	"zanko-khaledi/dsa/algorithmes"
)

type Edge = algorithmes.Edge

func main() {

	graph := map[string][]Edge{
		"A": {
			{To: "B", Weight: 4},
			{To: "C", Weight: 1},
			{To: "D", Weight: 2},
		},
		"B": {
			{To: "E", Weight: 5},
		},
		"C": {
			{To: "D", Weight: 1},
		},
		"D": {
			{To: "E", Weight: 3},
		},
		"E": {},
	}

	cost, path := algorithmes.DijkestraWithHeap(graph, "A")

	fmt.Println(cost)
	fmt.Println(path)
	fmt.Println(algorithmes.BuildPath(path, "A", "E"))
}
