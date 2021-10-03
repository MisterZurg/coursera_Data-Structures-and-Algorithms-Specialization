package main

import "fmt"

//  LeastCommonMultiple LCM(a,b) * GCD(a,b) = a*b
func LeastCommonMultiple(first, second int) int {
	return first * second / GCD(first, second)
}

// GCD is GCD_EuclideanAlgorithm from 1-3 Task
func GCD(first, second int) int {
	for first%second != 0 {
		first %= second
		first, second = second, first
	}
	return second
}

func main() {
	var first, second int
	fmt.Scan(&first, &second)
	// 6 8 = 24
	// 28851538 1183019 = 1933053046
	fmt.Println(LeastCommonMultiple(first, second))
}
