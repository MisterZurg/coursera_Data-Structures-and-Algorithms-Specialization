package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(sequence []int, l, r int) {
	if l >= r {
		return
	}
	m := Partition(sequence, l, r)

	QuickSort(sequence, l, m-1)
	QuickSort(sequence, m+1, r)
}

func Partition(sequence []int, l, r int) int {
	pivot := sequence[l]
	j := l
	for i := l + 1; i < r; i++ {
		j++
		// swap
		if sequence[i] > pivot {
			sequence[i], sequence[j] = sequence[j], sequence[i]
		}
	}
	sequence[l], sequence[j] = sequence[j], sequence[l]
	return j
}

func RandomizedQuickSort(sequence []int, l, r int) {
	if l >= r {
		return
	}
	// random number between l and r
	rand.Seed(time.Now().Unix())
	k := rand.Intn(r-l) + l
	sequence[l], sequence[k] = sequence[k], sequence[l]
	m := Partition(sequence, l, r)
	RandomizedQuickSort(sequence, l, m-1)
	RandomizedQuickSort(sequence, m+1, r)
}

// ImprovingQuickSort returns given sequence sorted in non-decreasing order.
func ImprovingQuickSort(sequence []int, l, r int) {
	// Here we have to implement 3-way partition
	if l >= r {
		return
	}
	m1, m2 := Partition3(sequence, l, r)

	ImprovingQuickSort(sequence, l, m1-1)
	ImprovingQuickSort(sequence, m2+1, r)
}

// Partition3 is a helper function for ImprovingQuickSort function
// realizes 3-way partition to handle few equal elements in slice
func Partition3(sequence []int, l, r int) (int, int) {
	pivot := sequence[l]
	m1 := l // We initiate m1 to be the part that is less than the pivot
	m2 := r // The part that is greater than the pivot
	i := l
	for i <= m2 {
		if sequence[i] < pivot {
			sequence[m1], sequence[i] = sequence[i], sequence[m1]
			m1++
			i++
		} else if sequence[i] > pivot {
			sequence[m2], sequence[i] = sequence[i], sequence[m2]
			m2--
		} else {
			i++
		}
	}

	return m1, m2
}

func main() {
	var n int
	var a []int

	fmt.Scan(&n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ImprovingQuickSort(a, 0, len(a)-1)
	// QuickSort(a, 0, len(a))
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}
}

/*
Input:
5
2 3 9 2 2
Output:
2 2 2 3 9
*/
