package main

import (
	"fmt"
)

// There's old Go compiler and no math.MaxInt
// const MAX_INT int = 9223372036854775807

// KINDA Is Graph Bipartite? - Leetcode 785
func isBipartite(graph map[int][]int, n int) bool {
	oddSet, evenSet := make(map[int]struct{}), make(map[int]struct{})
	visited := make([]bool, n)

	var helperBFS func(int, string) bool
	helperBFS = func(node int, cs string) bool {
		for _, adj_nd := range graph[node] {
			if !visited[adj_nd-1] {
				visited[adj_nd-1] = true

				switch cs {
				case "odd":
					evenSet[adj_nd] = struct{}{}
				case "even":
					oddSet[adj_nd] = struct{}{}
				}
			} else {
				if cs == "odd" {
					_, st1 := oddSet[node]
					_, st2 := oddSet[adj_nd]
					if st1 && st2 {
						return false
					}
				}
				if cs == "even" {
					_, st1 := evenSet[node]
					_, st2 := evenSet[adj_nd]
					if st1 && st2 {
						return false
					}
				}
			}

		}
		return true
	}
	currentSet := "odd"
	for v := 1; v <= n; v++ {
		if !visited[v-1] {
			visited[v-1] = true
		}
		token := helperBFS(v, currentSet)
		if !token {
			return false
		}

		if currentSet == "odd" {
			currentSet = "even"
		} else {
			currentSet = "odd"
		}
	}
	return true
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

	switch isBipartite(adj, n) {
	case true:
		fmt.Println(1)
	case false:
		fmt.Println(0)
	}
}
