package main

import "fmt"

// NumberofInversions counts the number of inversions of a given sequence.
func NumberofInversions(seq []int) int {
	_, inversions := MergeSort(seq)
	return inversions
}

// MergeSort sorts array seq and returns the number of inversions in seq.
func MergeSort(seq []int) ([]int, int) {
	if len(seq) == 1 {
		return seq, 0
	}
	m := len(seq) / 2

	seqLeft, leftNumOfPairs := MergeSort(seq[:m])
	seqRight, rightOfPairs := MergeSort(seq[m:])
	seq, number := Merge(seqLeft, seqRight)
	return seq, leftNumOfPairs + rightOfPairs + number
}

// Merge returns the resulting sorted array and the number of pairs
func Merge(seqLeft, seqRight []int) ([]int, int) {
	// seqLeft and seqRight are sorted
	// seqResult is empty slice of size seqLeft + seqRight
	seqResult := make([]int, len(seqLeft)+len(seqRight))
	resIdx := 0
	l := 0
	r := 0
	pairs := 0
	for l < len(seqLeft) && r < len(seqRight) {
		left := seqLeft[l]   // the first element of seqLeft
		right := seqRight[r] // the first element of seqRight
		if left <= right {
			// move left from seqLeft to the end of seqResult
			seqResult[resIdx] = left
			pairs++
			l++
		} else {
			// move right from seqRight to the end of seqResult
			seqResult[resIdx] = right
			r++
		}
		resIdx++
	}
	return seqResult, pairs
}

func main() {
	var n int
	var a []int

	fmt.Scan(&n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(NumberofInversions(a))
}

/*
Input:
5
2 3 9 2 9
Output:
2

Input:
4
1 2 3 4
Output:
0

Input:
6
9 8 7 3 2 1
Output:
15
*/
