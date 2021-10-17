package main

import "fmt"

func lastDigitOfFibSum(n int) int {
	if n <= 1 {
		return n
	}
	// Sum of nth Fibonacci series = F(n+2) -1

	// Then pisano period of module 10 = n+2 mod 60 = m
	// then find (F(m) mod 10) -1
	lesser := (n + 2) % 60

	if lesser == 1 {
		return 0
	} else if lesser == 0 {
		return 9
	}

	if lesser <= 1 {
		lesser -= 1
	} else {
		// Initialize two big ints with the first two numbers in the sequence.
		// F(n-1)
		var fn_m1 int = 0
		// F(n)
		var fn int = 1
		var sum int
		for i := 2; i < lesser+1; i++ {
			// Compute the next Fibonacci number, storing it in a.
			// Swap a and b so that b is the next number in the sequence.
			sum = (fn_m1 + fn) % 10
			fn, fn_m1 = sum, fn
		}
		lesser = sum - 1
	}
	if lesser < 0 {
		return lesser + 10
	}
	return lesser
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711
	var n int
	fmt.Scan(&n)
	fmt.Println(lastDigitOfFibSum(n))

	// fmt.Println(lastDigitOfFibSum(614162383528))
	// fmt.Println(lastDigitOfFibSum(3))
	// fmt.Println(lastDigitOfFibSum(239))
}

/*
Input:
3

Output:
4

Input:
239

Output:
0

Input:
614162383528

Output:
9
*/
