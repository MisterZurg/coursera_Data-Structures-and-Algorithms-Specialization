package main

import (
	"fmt"
)

// fibonacciNumberAgain возвращает отсаток от деления
// n-ого числа Фибоначчи по модулю modulo
func fibonacciNumberAgain(n, modulo int) int {
	period := getPisanoPeriod(modulo)
	n = n % period

	prev, cur := 0, 1

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		for i := 0; i < n-1; i++ {
			temp := cur
			cur = (prev + cur) % modulo
			prev = temp
		}
		return cur % modulo
	}
}

// В 1960 году Дональд Уолл из IBM, Phite Plains, Нью-Йорк,
// доказал, что множество чисел Фибоначчи, взятое по модулю m,
// является периодичным.

// getPisanoPeriod calculates and returns Pisano Period
// The length of a Pisano Period for
// a given m ranges from 3 to m * m
func getPisanoPeriod(modulo int) int {
	prev, cur, res := 0, 1, 0
	for i := 0; i < modulo*modulo; i++ {
		temp := 0
		temp = cur
		cur = (prev + cur) % modulo
		prev = temp

		if prev == 0 && cur == 1 {
			res = i + 1
		}
	}
	return res
}

func main() {
	var n, modulo int
	fmt.Scan(&n, &modulo)
	// 239 1000
	// F_239 mod 1 000 = 39 679 027 332 006 820 581 608 740 953 902 289 877 834 488 152 161 (mod 1 000) = 161.
	// 9999999999999 2
	fmt.Println(fibonacciNumberAgain(n, modulo))
}
