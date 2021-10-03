package main

import "fmt"

func GCD_Naive(first, second int) int {
	for first != second {
		if first > second {
			first -= second
		} else {
			second -= first
		}
	}
	return first
}

func GCD_EuclideanAlgorithm(first, second int) int {
	for first%second != 0 {
		first %= second
		first, second = second, first
	}
	return second
}

func main() {
	// 3918848
	// 1653264
	// = 61232
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(GCD_EuclideanAlgorithm(a, b))
}
