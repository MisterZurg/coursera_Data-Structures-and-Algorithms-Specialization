package main

import "fmt"

func MinRefils(stopsNumber, fuel int, stops []int) int {
	// Number of gas stations
	numRefils, currentRefil := 0, 0

	for currentRefil <= stopsNumber {
		lastRefil := currentRefil

		for currentRefil <= stopsNumber && stops[currentRefil+1]-stops[lastRefil] <= fuel {
			currentRefil++
		}

		if currentRefil == lastRefil {
			return -1
		}

		if currentRefil <= stopsNumber {
			numRefils++
		}
	}

	return numRefils
}

func main() {
	var distance, fuel, stopsNumber int
	fmt.Scan(&distance, &fuel, &stopsNumber)
	stops := make([]int, stopsNumber+2)
	stops[0], stops[len(stops)-1] = 0, distance
	for i := 1; i < len(stops)-1; i++ {
		fmt.Scan(&stops[i])
	}
	//fmt.Println(stops)

	fmt.Println(MinRefils(stopsNumber, fuel, stops))
}

/*
Input:
950
400
4
200 375 550 750
Output:
2

Input:
10
3
4
1 2 5 9
Output:
-1

*/
