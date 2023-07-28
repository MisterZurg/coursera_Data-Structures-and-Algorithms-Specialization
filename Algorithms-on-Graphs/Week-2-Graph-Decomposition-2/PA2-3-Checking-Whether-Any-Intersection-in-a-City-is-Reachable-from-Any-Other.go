package main

import (
	"fmt"
)

func dfs(u int, visited []bool, adj [][]int, stack *[]int) {
	visited[u] = true
	for _, v := range adj[u] {
		if !visited[v] {
			dfs(v, visited, adj, stack)
		}
	}
	*stack = append(*stack, u)
}

func reverse(adj [][]int) [][]int {
	n := len(adj)
	reversed := make([][]int, n)
	for u := 0; u < n; u++ {
		for _, v := range adj[u] {
			reversed[v] = append(reversed[v], u)
		}
	}
	return reversed
}

func numberOfStronglyConnectedComponents(adj [][]int) int {
	n := len(adj)
	visited := make([]bool, n)
	stack := make([]int, 0, n)

	// Perform DFS and store nodes in the order of their finish time in the stack
	for u := 0; u < n; u++ {
		if !visited[u] {
			dfs(u, visited, adj, &stack)
		}
	}

	// Reverse the graph
	reversedAdj := reverse(adj)

	// Perform DFS again, considering nodes in the order of their finish time (top of the stack)
	// and count the number of strongly connected components
	visited = make([]bool, n)
	stronglyConnected := make([][]int, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		u := stack[i]
		if !visited[u] {
			component := make([]int, 0)
			dfs(u, visited, reversedAdj, &component)
			stronglyConnected = append(stronglyConnected, component)
		}
	}

	return len(stronglyConnected)
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		adj[x-1] = append(adj[x-1], y-1)
	}

	fmt.Println(numberOfStronglyConnectedComponents(adj))
}

/*
5 7
2 1
3 2
3 1
4 3
4 1
5 2
5 3
*/
