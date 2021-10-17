package main

import "fmt"

func BinarySearch(sequence []int, target, low, high int) int {
	if high < low {
		return -1
	}

	mid := low + (high-low)/2
	if target == sequence[mid] {
		return mid
	} else if target < sequence[mid] {
		return BinarySearch(sequence, target, low, mid-1)
	} else {
		return BinarySearch(sequence, target, mid+1, high)
	}
}

func main() {
	var n, k int
	var a, b []int

	fmt.Scan(&n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Scan(&k)
	b = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&b[i])
	}
	// fmt.Println(a, b)

	for _, elem := range b {
		fmt.Printf("%d ", BinarySearch(a, elem, 0, len(a)-1))
	}
}
