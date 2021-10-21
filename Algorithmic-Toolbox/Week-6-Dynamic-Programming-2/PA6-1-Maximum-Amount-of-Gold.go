package main

import (
	"fmt"
)

// MaximumAmountOfGold implements Knapsack problem  without Repetitions
func MaximumAmountOfGold(W, n int, values, weights []int) int {
	K := make([][]int, n+1)
	for i := range K {
		K[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				K[i][w] = 0
			} else if weights[i-1] <= w {
				K[i][w] = Max(values[i-1]+K[i-1][w-weights[i-1]], K[i-1][w])
			} else {
				K[i][w] = K[i-1][w]
			}
		}
	}

	return K[n][W]
}

func Max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func main() {
	var W, n int
	fmt.Scan(&W, &n)
	// Weights of the bars of gold
	weights := make([]int, n)
	values := make([]int, n)

	for i := range weights {
		var num int
		fmt.Scan(&num)
		values[i] = num
		weights[i] = num
	}

	fmt.Println(MaximumAmountOfGold(W, n, values, weights))
}
