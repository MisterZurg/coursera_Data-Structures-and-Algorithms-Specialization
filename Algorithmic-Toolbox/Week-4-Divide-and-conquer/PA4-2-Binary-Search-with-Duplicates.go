package main

import "fmt"

func BinarySearchWithDuplicates(sequence []int, target int) int {
	// If the searched element located at index mid and its previous element
	// (i.e at index mid-1) match, binary search continues in the sorted space
	// to the left side of index mid i.e from index beg till index mid-1.
	// This way the search continues until the previous
	// element is the same as the element to be searched.
	// When the search terminates we get the index of the first occurrence.
	beg := 0
	end := len(sequence) - 1

	for beg <= end {

		mid := beg + (end-beg)/2

		if sequence[mid] == target {
			if mid-1 >= 0 && sequence[mid-1] == target {
				end = mid - 1
				continue
			}
			return mid
		} else if sequence[mid] < target {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
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
		fmt.Printf("%d ", BinarySearchWithDuplicates(a, elem))
	}
}

/*
7
2 4 4 4 7 7 9
4
9 4 5 2

Correct output:
6 1 -1 0
*/
