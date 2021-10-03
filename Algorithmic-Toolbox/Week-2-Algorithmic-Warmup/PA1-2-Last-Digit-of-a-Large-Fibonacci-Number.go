package main

import (
	"fmt"
	// "math/big"
)

func lastDigitOfFibonacciNumber(n int) uint { // *big.Int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		var fn, fn_m1 uint = 1, 1
		for i := 2; i < n; i++ {
			fn_m1, fn = fn, (fn+fn_m1)%10
		}
		return fn
	}
}

func main() {
	// 1, 1, 2, 3, 5, 8, 13, 21
	// F_8 = 3
	// F_9 = 1
	// F_13 = 144
	// F_18 = 1597
	// 280 571 172 992 510 140 037 611 932 413 038 677 189 525
	// F_200 = ... 5
	// F_331 = ... 9
	// F_327305 = ... 5
	// F_613455 = ...
	var n int
	fmt.Scan(&n)
	fmt.Println(lastDigitOfFibonacciNumber(n))
}
