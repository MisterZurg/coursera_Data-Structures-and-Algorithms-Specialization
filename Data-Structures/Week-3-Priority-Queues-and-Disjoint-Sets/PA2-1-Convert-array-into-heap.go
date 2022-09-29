package main

import (
	"fmt"
	"math"
)

func main() {

	// Test testcases :)
	// swaps_slice := MakeHeapFromArray([]int{5, 4, 3, 2, 1})
	// swaps_slice := MakeHeapFromArray([]int{1, 2, 3, 4,5})

	inputSlice := ParseInput()
	swapsSlice := MakeHeapFromArray(inputSlice)
	PrintAnswer(swapsSlice)
}

func ParseInput() []int {
	// Input Format.
	// The first line of the input contains single integer n.
	var n int
	fmt.Scan(&n)
	// The next line contains n space-separated integers a_i.
	arrayButActuallySlice := make([]int, n)
	for elem := range arrayButActuallySlice {
		fmt.Scan(&arrayButActuallySlice[elem])
	}

	return arrayButActuallySlice
}

func PrintAnswer(swaps []string) {
	fmt.Println(len(swaps))
	for _, swapPair := range swaps {
		fmt.Println(swapPair)
	}
}

// MakeHeapFromArray takes an array as an input and makes min heap
func MakeHeapFromArray(unsorted []int) []string {
	swaps := make([]string, 0)
	n := len(unsorted)
	for i := (n - 1) / 2; i+1 != 0; i-- {
		SiftDown(i, unsorted, &swaps)
	}
	return swaps
}

func SiftDown(idx int, H []int, swaps *[]string) {
	minIndex := idx
	l := LeftChild(idx)
	size := len(H)

	if l < size && H[l] < H[minIndex] {
		minIndex = l
	}

	r := RightChild(idx)
	if r < size && H[r] < H[minIndex] {
		minIndex = r
	}

	if idx != minIndex {
		H[idx], H[minIndex] = H[minIndex], H[idx]

		swap_pair := fmt.Sprintf("%d %d", idx, minIndex)
		*swaps = append(*swaps, swap_pair)

		SiftDown(minIndex, H, swaps)
	}
}

func LeftChild(idx int) int {
	return 2*idx + 1
}

func RightChild(idx int) int {
	return 2*idx + 2
}

func Parent(idx int) int {
	return int(math.Floor(float64(idx-1) / 2))
}
