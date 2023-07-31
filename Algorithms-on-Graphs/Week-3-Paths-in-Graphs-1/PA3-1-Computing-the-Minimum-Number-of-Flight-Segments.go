package main

import (
	"fmt"
)

// There's old Go compiler and no math.MaxInt
const MAX_INT int = 9223372036854775807

func BFS(graph map[int][]int, n, a, b int) int {
	dist := make([]int, n)
	prev := make([]int, n)
	for i := range dist {

		dist[i] = MAX_INT
		prev[i] = -1 // nil
	}
	dist[a-1] = 0

	q := []int{a}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		for _, v := range graph[u] {
			if dist[v-1] == MAX_INT {
				q = append(q, v)
				dist[v-1] = dist[u-1] + 1
				prev[v-1] = u - 1
			}
		}
	}
	if dist[b-1] == MAX_INT {
		return -1
	}
	return dist[b-1]
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	adj := make(map[int][]int)
	var v, u int
	for i := 0; i < m; i++ {
		fmt.Scan(&v, &u)
		adj[v] = append(adj[v], u)
		adj[u] = append(adj[u], v)
	}
	fmt.Scan(&v, &u)
	fmt.Println(BFS(adj, n, v, u))
}
