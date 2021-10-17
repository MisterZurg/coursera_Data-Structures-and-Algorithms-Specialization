package main

import (
	"fmt"
	"math"
)

const (
	TEN  = 10
	FIVE = 5
)

func MoneyChange(m int) int {
	changeCounter := 0
	changeCounter += int(math.Floor(float64(m / TEN)))
	m %= TEN
	changeCounter += int(math.Floor(float64(m / FIVE)))
	m %= FIVE
	changeCounter += m
	return changeCounter
}

func main() {
	var price int
	fmt.Scan(&price)
	fmt.Println(MoneyChange(price))
	// fmt.Println(MoneyChange(28))
	// fmt.Println(MoneyChange(2))
}
