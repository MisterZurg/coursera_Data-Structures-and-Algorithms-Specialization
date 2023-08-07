package main

import (
	"fmt"
	"math"
)

const (
	INF           = 1<<63 - 1
	INF_BUT_FLOAT = 0x1p1023 * (1 + (1 - 0x1p-52))
)

type Point struct {
	x int
	y int
}

func LengthOfASegment(first, second Point) float64 {
	xs := math.Pow(float64(first.x-second.x), 2)
	ys := math.Pow(float64(first.y-second.y), 2)
	return math.Sqrt(xs + ys)
}

func findMin(set map[int]struct{}, keys *[]float64) int {
	minID := INF
	minFloat := INF_BUT_FLOAT

	for ID := range set {
		if (*keys)[ID] < minFloat {
			minID = ID
			minFloat = (*keys)[ID]
		}
	}

	return minID
}

func minimumDistance(points []Point) (minimalDistance float64) {
	n := len(points)
	// Construct 2D matrix for distance
	dist := make([][]float64, n)
	for i := range dist {
		dist[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := LengthOfASegment(points[i], points[j])
			dist[i][j] = d
			dist[j][i] = d
		}
	}
	keys := make([]float64, n)
	for i := range keys {
		keys[i] = INF_BUT_FLOAT
	}
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	keys[0] = 0
	Q := make(map[int]struct{})
	for i := 0; i < n; i++ {
		Q[i] = struct{}{}
	}

	for len(Q) > 0 {
		// Find the min(Q) by key value
		u := findMin(Q, &keys)
		delete(Q, u)
		if parent[u] != -1 {
			minimalDistance += dist[u][parent[u]]
		}

		for v := 0; v < n; v++ {
			if dist[u][v] != 0 {
				_, ok := Q[v]

				if ok && dist[u][v] < keys[v] {
					parent[v] = u
					keys[v] = dist[u][v]
				}
			}
		}
	}

	return minimalDistance
}

func main() {
	var n int
	fmt.Scan(&n)
	points := make([]Point, n)
	for i := range points {
		fmt.Scan(&points[i].x, &points[i].y)
	}

	fmt.Println(minimumDistance(points))
}
