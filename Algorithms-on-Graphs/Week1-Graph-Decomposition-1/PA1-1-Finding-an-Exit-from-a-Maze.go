package main

import "fmt"

type adjacencyList map[int][]int

func parseInput() ([]int, adjacencyList) {
	// edgeList -> adjList
	var n, m int
	fmt.Scan(&n, &m)

	// 1 : 2, 3, 4
	// 2 : 1, 4, 6
	// ...
	al := make(adjacencyList)

	for i := 0; i < m; i++ {
		var v, u int
		fmt.Scan(&v, &u)
		al[v] = append(al[v], u)
		al[u] = append(al[u], v)
	}

	var v1, v2 int
	fmt.Scan(&v1, &v2)
	return []int{v1, v2}, al
}

func DFS(path []int, G adjacencyList) {
	visited := make(map[int]bool)

	var Explore func(int)
	Explore = func(v int) {
		visited[v] = true
		for _, w := range G[v] {
			if !visited[w] {
				Explore(w)
			}
		}
	}

	Explore(path[0])
	if visited[path[1]] {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func main() {
	path, al := parseInput()
	DFS(path, al)
}
