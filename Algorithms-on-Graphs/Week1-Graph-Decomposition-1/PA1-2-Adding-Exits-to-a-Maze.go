package main

import "fmt"

type adjacencyList map[int][]int

func parseEdgeList() adjacencyList {
	// edgeList -> adjList
	var n, m int
	fmt.Scan(&n, &m)

	// 1 : 2, 3, 4
	// 2 : 1, 4, 6
	// ...
	al := make(adjacencyList)
	for v := 1; v <= n; v++ {
		al[v] = nil
	}

	for i := 0; i < m; i++ {
		var v, u int
		fmt.Scan(&v, &u)
		al[v] = append(al[v], u)
		al[u] = append(al[u], v)
	}

	return al
}

func DFSCountComponents(G adjacencyList) {
	visited := make(map[int]bool)
	// connected component
	cc := 0

	var Explore func(int)
	Explore = func(v int) {
		visited[v] = true

		// In case we have to save what connected component number
		// is that vertex
		// CCnum(v) ← cc
		for _, w := range G[v] {
			if !visited[w] {
				Explore(w)
			}
		}
	}

	// Traverse through graph vertices
	for v := range G {
		// Default value for []int is zero so we avoid going into non-existed 0 vertex
		if G[v] == nil && !visited[v] {
			visited[v] = true
			cc++
			continue
		}
		if !visited[v] {
			Explore(v)
			cc++
		}
	}
	fmt.Println(cc)
}

func main() {
	al := parseEdgeList()
	DFSCountComponents(al)
}

/*
Test Case # 2
Input:
4 3
1 2
3 2
4 3

Correct output:
1
*/
