package main

import (
	"fmt"
	"sort"
)

func MaximumAdvertisementRevenue(profit, clicks []int) int {
	stonks := 0
	sort.Ints(profit)
	sort.Ints(clicks)
	for idx := range profit {
		stonks += profit[idx] * clicks[idx]
	}
	return stonks
}

func main() {
	var n int
	var profitPerClick, clicksPerDay []int
	fmt.Scan(&n)
	// Fill profitPerClick
	var elem int
	for i := 0; i < n; i++ {
		fmt.Scan(&elem)
		profitPerClick = append(profitPerClick, elem)
	}
	// Clicks
	for i := 0; i < n; i++ {
		fmt.Scan(&elem)
		clicksPerDay = append(clicksPerDay, elem)
	}
	fmt.Println(MaximumAdvertisementRevenue(profitPerClick, clicksPerDay))

}
