package main

import "fmt"

func lastDigitOfFibSumOfSquares(n int) int {
	fn_m1 := n % 60
	fn := (n + 1) % 60
	if fn_m1 > 1 {
		fn_m1 = SumOfSquares(fn_m1)
	}

	if fn > 1 {
		fn = SumOfSquares(fn)
	}
	return (fn_m1 * fn) % 10
}

func SumOfSquares(n int) int {
	// Initialize two big ints with the first two numbers in the sequence.
	// F(n-1)
	var fn_m1 int = 0
	// F(n)
	var fn int = 1
	var sum int
	for i := 1; i < n+1; i++ {
		// Compute the next Fibonacci number, storing it in a.
		// Swap a and b so that b is the next number in the sequence.
		sum = (fn_m1 + fn) % 10
		fn_m1, fn = sum, fn_m1
	}
	return sum
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711
	var n int
	fmt.Scan(&n)
	fmt.Println(lastDigitOfFibSumOfSquares(n))
}
