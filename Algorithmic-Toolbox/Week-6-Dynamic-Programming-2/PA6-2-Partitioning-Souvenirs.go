package main

import (
	"fmt"
	"sort"
)

// Partition To K Equal Sum Subsets From An Array of Integers
func PartitioningSouvenirs(nums []int, k int) int {
	var sum = 0
	for i := range nums {
		sum += nums[i]
	}
	// if sum of all numbers is not divisible for k, there is something wrong with input data
	if sum%k != 0 {
		return 0
	}
	// target is the sum for each subsets
	target := sum / k
	// sort by desc to filtrate some values
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var i int
	// filtrate for nums which equal target, as number '5' in Example 1
	for i < len(nums) && nums[i] == target {
		i++
		k--
	}
	// if there is no subsets left, so we placed all numbers. hurray!
	if k == 0 {
		return 1
	}
	// subsets is the slice, we should place our numbers till every subset be full (subsets[i] == target)
	subsets := make([]int, k, k)

	return put(nums[i:], subsets, target)
}

func put(nums []int, subsets []int, target int) int {
	// base case, if there is left no numbers, it mean everything is placed
	if len(nums) == 0 {
		return 1
	}
	// iterate through subsets and try to put number (nums[0]) to each subset
	for i := range subsets {
		if subsets[i]+nums[0] <= target {
			// if subsets[i] is not fuul and can handle number nums[0]: add num[0] to subset
			subsets[i] += nums[0]
			// check if we could place next number in any subset
			if put(nums[1:], subsets, target) == 0 {
				// if not - backtracking
				subsets[i] -= nums[0]
			} else {
				// everything ok
				return 1
			}
		}
	}
	// there is no any valid allocation
	return 0
}

func main() {
	var n int
	fmt.Scan(&n)
	items := make([]int, n)
	for i := range items {
		fmt.Scan(&items[i])
	}

	fmt.Println(PartitioningSouvenirs(items, 3))
}
