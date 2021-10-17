package main

import (
	"fmt"
	"sort"
)

func MaximumValueoftheLoot(n, capacity float64) float64 {
	var fromMax2Min []float64
	var val, wgh float64
	prices := make(map[float64]float64)
	for i := 0; i < int(n); i++ {
		fmt.Scan(&val, &wgh)
		fromMax2Min = append(fromMax2Min, val/wgh)
		prices[val/wgh] = wgh
	}
	// Вычитать если можно иначе делитиь и вычитать
	sort.Float64s(fromMax2Min)
	var maximumValue float64 = 0
	for i := len(fromMax2Min) - 1; i >= 0; i-- {
		if capacity >= prices[fromMax2Min[i]] {
			capacity -= prices[fromMax2Min[i]]
			maximumValue += fromMax2Min[i] * prices[fromMax2Min[i]]
		} else if 0 < capacity {
			//if capacity < prices[fromMax2Min[i]] {
			maximumValue += fromMax2Min[i] * capacity
			return maximumValue
		}
		// maximumValue += fromMax2Min[i] * capacity
	}
	return maximumValue
}

func main() {
	var n, capacity float64
	fmt.Scan(&n, &capacity)
	fmt.Printf("%.4f", MaximumValueoftheLoot(n, capacity))
}
