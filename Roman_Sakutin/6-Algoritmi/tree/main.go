package main

import "fmt"

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 4, 5},
		3: {1, 5, 6},
		4: {2, 5},
		5: {2, 3, 4, 6},
		6: {3, 5, 7},
		7: {6, 8, 9},
		8: {7, 9},
		9: {7, 8},
	}
	fmt.Println(bfs(graph))
}
