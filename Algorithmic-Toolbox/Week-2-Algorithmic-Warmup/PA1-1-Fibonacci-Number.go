package main

import "fmt"

func fibonacciNumber(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		var memorizedFib []int = make([]int, n)
		memorizedFib[0], memorizedFib[1] = 1, 1
		for i := 2; i < n; i++ {
			memorizedFib[i] = memorizedFib[i-1] + memorizedFib[i-2]
		}

		return memorizedFib[n-1]
	}
}

func fibbonaciNumberWitgoutMemoriization(n int) int {
	// F(n)
	var x = 1
	// F(n-1)
	var y = 0
	for i := 0; i < n; i++ {
		x += y
		y = x - y
	}
	return y
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacciNumber(n))
}
