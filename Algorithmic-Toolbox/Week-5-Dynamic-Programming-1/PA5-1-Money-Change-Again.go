package main

import (
	"fmt"
)

// Money Change Problem for denominations 1, 3, and 4.
func MoneyChangeAgain(money int) int {
	coins := []int{1, 3, 4}

	minNumCoins := make([]int, money+1)
	for i := 0; i < len(minNumCoins); i++ {
		// As it said in lecture we should
		// fill minNumCoins with inf values
		minNumCoins[i] = 9223372036854775807 // math.MaxInt
	}
	minNumCoins[0] = 0
	var NumCoins int
	for m := 1; m <= money; m++ {
		// What is the minimum number of Coins needed to change
		// m cents for denominations 1, 3, and 4.
		for _, coin := range coins {
			if m >= coin {
				NumCoins = minNumCoins[m-coin] + 1
				// fmt.Println(NumCoins)
				if NumCoins < minNumCoins[m] {
					minNumCoins[m] = NumCoins
				}
			}
		}
	}
	return minNumCoins[money]
}

func main() {
	var money int
	fmt.Scan(&money)
	fmt.Println(MoneyChangeAgain(money))
}

/*
Input:
2
Output:
2


Input:
34
Output:
9
*/
