package main

import (
	"fmt"
	"strconv"
)

func MaximumSalary(numbers []int) string {
	salary := ""

	for len(numbers) > 0 {
		maxDigit, maxInd := LargestNumber(numbers)
		salary += strconv.Itoa(maxDigit)
		numbers = RemoveIndex(numbers, maxInd)
	}
	return salary
}

func LargestNumber(numbers []int) (int, int) {
	maxDigit := 0
	maxInd := 0
	for idx, number := range numbers {
		if IsGreaterOrEqual(number, maxDigit) {
			maxDigit = number
			maxInd = idx
		}
	}
	return maxDigit, maxInd
}

func IsGreaterOrEqual(digit, maxDigit int) bool {
	first := strconv.Itoa(digit) + strconv.Itoa(maxDigit)
	firstNum, _ := strconv.Atoi(first)

	second := strconv.Itoa(maxDigit) + strconv.Itoa(digit)
	secondNum, _ := strconv.Atoi(second)

	return firstNum >= secondNum
}

func RemoveIndex(s []int, index int) []int {
	// Remove the element at index i from a.
	s[index] = s[len(s)-1] // Copy last element to index i.
	s[len(s)-1] = 0        // Erase last element (write zero value).
	s = s[:len(s)-1]       // Truncate slice.
	return s
}

func main() {
	var number4Salary int
	fmt.Scan(&number4Salary)

	numbers := make([]int, number4Salary)

	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	fmt.Println(MaximumSalary(numbers))
}
