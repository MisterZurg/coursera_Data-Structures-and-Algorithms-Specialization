package main

import "fmt"

func lastDigitOfFibonacciNumberAgain(from, to int) int {
	if to == 0 {
		return 0
	}

	lesser_m := (from + 1) % 60
	lesser_n := (to + 2) % 60

	if lesser_n <= 1 {
		lesser_n -= 1
	} else {
		lesser_n = SumS(lesser_n)
	}

	if lesser_m <= 1 {
		lesser_m -= 1
	} else {
		lesser_m = SumS(lesser_m)
	}

	if lesser_n >= lesser_m {
		return (lesser_n - lesser_m)
	}
	return 10 + lesser_n - lesser_m
}

func SumS(n int) int {
	// Initialize two big ints with the first two numbers in the sequence.
	// F(n-1)
	var fn_m1 int = 0
	// F(n)
	var fn int = 1
	var sum int
	for i := 2; i < n+1; i++ {
		// Compute the next Fibonacci number, storing it in a.
		// Swap a and b so that b is the next number in the sequence.
		sum = (fn_m1 + fn) % 10
		fn, fn_m1 = sum, fn
	}
	return sum - 1
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711
	var from, to int
	fmt.Scan(&from, &to)
	fmt.Println(lastDigitOfFibonacciNumberAgain(from, to))
}

/*
Input:
0 7

Correct output:
3
*/
