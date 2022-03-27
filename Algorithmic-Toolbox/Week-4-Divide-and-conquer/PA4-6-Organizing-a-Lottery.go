package main

import (
	"fmt"
)

/*
Cash on me, like I hit the lottery (the lottery)
#@$ will trip, watch them how they follow me (wait)
Hunnids blue, yeah I got them all on me
Go, go, go, go, let's go
Prada shoes, yeah I keep a style on me (style on me)
Pretty freaks, make them #@$ pile on me (I swear)
Rack party, I got thirty thou' on me (right now)
Go, go, go, go, let's go
*/

func OrganizingALottery(points []int, segments [][]int) []int {
	answer := make([]int, len(points))
	for ind, point := range points {
		for _, segment := range segments {
			if segment[0] <= point && point <= segment[1] {
				answer[ind]++
			}
		}
	}
	return answer
}

func main() {
	var numberOfSegments, numberOfPoints int
	fmt.Scan(&numberOfSegments, &numberOfPoints)

	segments := make([][]int, numberOfSegments)
	for i := range segments {
		segments[i] = make([]int, 2)
	}

	for i := range segments {
		for j := range segments[i] {
			fmt.Scan(&segments[i][j])
		}
	}

	points := make([]int, numberOfPoints)
	for i := range points {
		fmt.Scan(&points[i])
	}

	answer := OrganizingALottery(points, segments)

	for _, k := range answer {
		fmt.Printf("%d ", k)
	}
}
