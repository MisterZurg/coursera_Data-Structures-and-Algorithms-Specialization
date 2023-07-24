package main

import "fmt"

type adjacencyList map[int][]int

// Ref: "Detect Cycle in Directed Graph" URL https://www.scaler.com/topics/detect-cycle-in-directed-graph/
func detectCycleInDAG(DAG adjacencyList) bool {
	visited := make([]bool, len(DAG)+1)
	inStack := make([]bool, len(DAG)+1)

	var DFS func(int, []bool, []bool) bool
	DFS = func(node int, visited, inStack []bool) bool {
		// Check if node exists in the
		// recursive stack.
		if inStack[node] {
			return true
		}

		// Check if node is already visited.
		if visited[node] {
			return false
		}
		// Marking node as visited.
		visited[node] = true
		// Marking node to be present in
		// recursive stack.
		inStack[node] = true

		// Iterate for all adjacent of 'node'.
		for _, v := range DAG[node] {
			// fmt.Println("curr=", v)
			// Recurse for 'v'.
			if DFS(v, visited, inStack) {
				return true
			}
		}

		// Mark 'node' to be removed
		// from the recursive stack.
		inStack[node] = false

		// Return false if no cycle exists.
		return false
	}

	for vertex := range DAG {
		// fmt.Println(vertex)
		// Check if cycle exists.
		if DFS(vertex, visited, inStack) {
			return true
		}
	}
	// Returning false, if no cycle is found.
	return false
}

func ScanDAG() adjacencyList {
	var n, m int     // vertices,  edges
	fmt.Scan(&n, &m) // 1 ≤ n ≤ 10^3, 0 ≤ n ≤ 10^3.

	adj := make(adjacencyList)
	for i := 1; i <= n; i++ {
		adj[i] = nil
	}

	// Scanning directed graph
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u] = append(adj[u], v)
	}
	return adj
}

func main() {
	DAG := ScanDAG()
	// fmt.Println(DAG)
	if detectCycleInDAG(DAG) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
