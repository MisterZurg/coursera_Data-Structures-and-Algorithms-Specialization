package main

import (
	"fmt"
)

const INF int = 9223372036854775807

func shortestPaths(adj [][]int, cost [][]int, s int, distance []int, reachable []int, shortest []int) {
	n := len(adj)
	distance[s] = 0
	reachable[s] = 1

	for i := 0; i <= n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < len(adj[j]); k++ {
				if reachable[j] == 1 && distance[adj[j][k]] > distance[j]+cost[j][k] {
					distance[adj[j][k]] = distance[j] + cost[j][k]
					reachable[adj[j][k]] = 1
					if i == n {
						shortest[adj[j][k]] = 0
					}
				}
			}
		}
	}

	queue := make([]int, 0)
	for i, v := range shortest {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		top := queue[0]
		shortest[top] = 0
		for i := 0; i < len(adj[top]); i++ {
			if shortest[adj[top][i]] == 1 {
				queue = append(queue, adj[top][i])
			}
		}
		queue = queue[1:]
	}
}

func main() {
	var n, m, s int
	fmt.Scan(&n, &m)
	adj := make([][]int, n)
	cost := make([][]int, n)
	for i := 0; i < m; i++ {
		var x, y, w int
		fmt.Scan(&x, &y, &w)
		adj[x-1] = append(adj[x-1], y-1)
		cost[x-1] = append(cost[x-1], w)
	}
	fmt.Scan(&s)
	s--
	distance := make([]int, n)
	reachable := make([]int, n)
	shortest := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = INF
		reachable[i] = 0
		shortest[i] = 1
	}
	shortestPaths(adj, cost, s, distance, reachable, shortest)
	for i := 0; i < n; i++ {
		if reachable[i] == 0 {
			fmt.Println("*")
		} else if shortest[i] == 0 {
			fmt.Println("-")
		} else {
			fmt.Println(distance[i])
		}
	}
}
