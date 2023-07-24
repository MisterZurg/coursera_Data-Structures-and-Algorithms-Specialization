package main

import "fmt"

type G struct {
	Adj AdjacencyList
	V   int // Vertex number
}
type AdjacencyList map[int][]int

func ScanDAG() G {
	var n, m int     // vertices,  edges
	fmt.Scan(&n, &m) // 1 ≤ n ≤ 10^3, 0 ≤ n ≤ 10^3.

	adj := make(AdjacencyList)
	for i := 1; i <= n; i++ {
		adj[i] = nil
	}

	// Scanning directed graph
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u] = append(adj[u], v)
	}
	return G{
		Adj: adj,
		V:   n,
	}
}

func TopologicalSort(graph G) {
	visited := make([]bool, graph.V+1) // Vertex from 1 ...
	var stack []int

	var DFS func(node int)
	DFS = func(node int) {
		if visited[node] {
			return
		}

		visited[node] = true

		for _, v := range graph.Adj[node] {
			if !visited[v] {
				DFS(v)
			}
		}

		stack = append(stack, node)
		// return true
	}

	for i := 1; i <= graph.V; i++ {
		DFS(i)
	}
	// sort reverse postorder
	// reverse print stack
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("%d ", stack[i])
	}
}

func main() {
	TopologicalSort(ScanDAG())
}

/*
4 3
1 2
4 1
3 1
Output:
4 3 1 2

Input:
4 1
3 1
Output:
2 3 1 4

Input:
5 7
2 1
3 2
3 1
4 3
4 1
5 2
5 3
Output:
5 4 3 2 1
*/
