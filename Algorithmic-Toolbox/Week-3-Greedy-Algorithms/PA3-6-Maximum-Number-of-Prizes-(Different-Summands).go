package main

import "fmt"

func MaximumNumberOfPrizes(sum int) {
	values := getValues(sum)
	fmt.Println(len(values))
	for _, val := range values {
		fmt.Printf("%d ", val)
	}
}

func getValues(sum int) []int {
	var values []int
	var value int

	for i := 1; sum > 0; i++ {

		if sum <= (2 * i) {
			value = sum
		} else {
			value = i
		}

		values = append(values, value)
		sum -= value
	}
	return values
}

func main() {
	var number int
	fmt.Scan(&number)
	MaximumNumberOfPrizes(number)
}
