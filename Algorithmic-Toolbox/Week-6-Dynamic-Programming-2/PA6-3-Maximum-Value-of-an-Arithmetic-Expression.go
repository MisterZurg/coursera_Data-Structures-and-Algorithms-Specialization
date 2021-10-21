package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// Prentheses is computes and caches the result
func Prentheses(numbers []int, operations []string) int {
	n := len(numbers)

	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	M := make([][]int, n)
	for i := range m {
		M[i] = make([]int, n)
	}

	// Like it' said in lecture
	for i := 0; i < n; i++ {
		m[i][i] = numbers[i]
		M[i][i] = numbers[i]
	}

	// Main logic
	for s := 1; s < n; s++ {
		for i := 0; i < n-s; i++ {
			j := i + s
			M[i][j], m[i][j] = MaxAndMin(M, m, i, j, operations)
		}
	}

	return M[0][n-1]
}

// MaxAndMin is our core function to full-fill the answer table
func MaxAndMin(M, m [][]int, i, j int, operations []string) (int, int) {
	// min_value = +inf and max_value = -inf
	min_value := 9223372036854775807  // math.MaxInt
	max_value := -9223372036854775808 // math.MaxInt

	for k := i; k < j; k++ {
		a := Calc(M[i][k], M[k+1][j], operations[k])
		b := Calc(M[i][k], m[k+1][j], operations[k])
		c := Calc(m[i][k], M[k+1][j], operations[k])
		d := Calc(m[i][k], m[k+1][j], operations[k])

		max_value = MaxOfTwo(max_value, MaxOfTwo(a, MaxOfTwo(b, MaxOfTwo(c, d))))
		min_value = MinOfTwo(min_value, MinOfTwo(a, MinOfTwo(b, MinOfTwo(c, d))))
	}
	return max_value, min_value
}

// MaxOfTwo is a helper func
func MaxOfTwo(first, second int) int {
	if first > second {
		return first
	}
	return second
}

// MinOfTwo is a helper func
func MinOfTwo(first, second int) int {
	if first < second {
		return first
	}
	return second
}

// Calc is a helper function that evaluates the expression
// from two numbers and operation between them
func Calc(first, second int, op string) int {
	switch op {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	default:
		return 0
	}
}

func main() {
	var expression string
	fmt.Scan(&expression)

	numbers := make([]int, 0)
	operations := make([]string, 0)
	// Split expression into numbers and operation slice
	for _, symbol := range expression {
		if unicode.IsDigit(symbol) {
			number, _ := strconv.Atoi(string(symbol))
			numbers = append(numbers, number)
		} else {
			operations = append(operations, string(symbol))
		}
	}
	fmt.Println(Prentheses(numbers, operations))
}

/*
	Input:
	1+5

	Output:
	6


	Input:
	5-8+7*4-8+9

	Output:
	200
*/
