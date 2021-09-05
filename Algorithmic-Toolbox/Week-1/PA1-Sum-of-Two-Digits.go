/*
In this very first programming challenge, your goal is to implement a program
that reads two digits from the standard input and prints their sum to the standard output.
We start from this simple problem to show you the pipeline of submitting a solution to the grading system.
 */
package Week_1

import "fmt"

func SumOfTwoDigits() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a + b)
}