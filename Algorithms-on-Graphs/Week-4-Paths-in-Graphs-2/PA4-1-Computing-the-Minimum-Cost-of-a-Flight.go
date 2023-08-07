//// Kinda 787. Cheapest Flights Within K Stops
//// Also check
//// "Implementation of Dijkstra using heap in Go" by Douglas Makey Mendez Molero https://dev.to/douglasmakey/implementation-of-dijkstra-using-heap-in-go-6e3
package main

import (
	"container/heap"
	"fmt"
)

type Vertex struct {
	index int
	dist  int
}

type PriorityQueue []Vertex

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Vertex)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

//func distance(adj [][]int, cost [][]int, s, t int) int {
func (g *WeightedDAG) Dijkstra(s, t int) int {
	const INF int = 9223372036854775807
	dist := make([]int, g.nodesNum)
	for i := range dist {
		dist[i] = INF
	}
	dist[s] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Vertex{s, 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(Vertex).index

		if u == t {
			break
		}

		for _, v := range (*g).adj[u] {
			if dist[u]+v.Cost < dist[v.toNode] {
				dist[v.toNode] = dist[u] + v.Cost
				heap.Push(pq, Vertex{v.toNode, dist[v.toNode]})
			}
		}
	}

	if dist[t] == INF {
		return -1
	}

	return dist[t]
}

type WeightedDAG struct {
	adj      map[int][]WeightedEdge
	nodesNum int
}

func newWeightedDAG(n int) WeightedDAG {
	return WeightedDAG{
		adj:      make(map[int][]WeightedEdge),
		nodesNum: n,
	}
}

type WeightedEdge struct {
	toNode int
	Cost   int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	//adj := make([][]int, n)
	//cost := make([][]int, n)
	g := newWeightedDAG(n)
	for i := 0; i < m; i++ {
		var x, y, w int
		fmt.Scan(&x, &y, &w)

		g.adj[x-1] = append(g.adj[x-1], WeightedEdge{toNode: y - 1, Cost: w})
		//adj[x-1] = append(adj[x-1], y-1)
		//cost[x-1] = append(cost[x-1], w)
	}

	var s, t int
	fmt.Scan(&s, &t)
	//fmt.Println(distance(adj, cost, s-1, t-1))
	fmt.Println(g.Dijkstra(s-1, t-1))
}

/*
4 4
1 3 1
1 2 1
2 4 4
3 4 2
1 4

*/
