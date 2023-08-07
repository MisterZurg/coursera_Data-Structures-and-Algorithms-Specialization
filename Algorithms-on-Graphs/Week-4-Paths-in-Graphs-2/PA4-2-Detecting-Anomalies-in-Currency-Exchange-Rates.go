package main

import (
	"fmt"
)

func BellmanFord(adj [][]int, cost [][]int) int {
	const INF int = 2147483647
	n := len(adj)
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = INF
	}

	dist[0] = 0

	for i := 0; i <= n; i++ {
		for u := 0; u < n; u++ {
			for k := 0; k < len(adj[u]); k++ {
				v := adj[u][k]
				if dist[v] > dist[u]+cost[u][k] {
					dist[v] = dist[u] + cost[u][k]
					if i == n {
						return 1
					}
				}
			}
		}
	}

	return 0
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	adj := make([][]int, n)
	cost := make([][]int, n)

	for i := 0; i < m; i++ {
		var x, y, w int
		fmt.Scan(&x, &y, &w)
		adj[x-1] = append(adj[x-1], y-1)
		cost[x-1] = append(cost[x-1], w)
	}

	fmt.Println(BellmanFord(adj, cost))
}
