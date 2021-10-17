package main

import "fmt"

func CarFueling(distance, m, stops int) int {
	// Create path []slice
	fillingStation := make([]int, stops+2) // + 2 cause Start = 0 and end = distance
	// Set start and end point
	fillingStation[0], fillingStation[stops+1] = 0, distance
	for i := 1; i < stops+1; i++ {
		fmt.Scan(&fillingStation[i])
	}
	currentFuel := m
	stopsCount := 0
	for i := 1; i < len(fillingStation)-1; i++ {
		currentDistance := fillingStation[i] - fillingStation[i-1]
		//currentNextDistance := fillingStation[i + 1] - fillingStation[i]
		if currentFuel-currentDistance >= 0 {
			currentFuel -= currentDistance
			if currentFuel-currentDistance-fillingStation[i-1] >= 0 {
				stopsCount--
			} else {
				stopsCount++
				currentFuel = m
			}
		} else {
			return -1
		}
	}
	if fillingStation[len(fillingStation)-1]-fillingStation[len(fillingStation)-2] > currentFuel {
		return -1
	}
	return stopsCount
}

func main() {
	var distance, fuel, stops int
	fmt.Scan(&distance, &fuel, &stops)
	fmt.Println(CarFueling(distance, fuel, stops))
	// Input:
	// 		950
	// 		400
	// 		4
	// 		200 375 550 750
	// Output:
	// 		2
	//
	//	700
	//	200
	//	4
	//	100 200 300 400
	// -1
}
