package main

import "fmt"

// Majority rule is a decision rule that selects
// the alternative which has a majority,
// that is, more than half the votes.
// Time Complexity: O(n)
func MajorityElement(votes []int) int {
	elections := make(map[int]int)
	for _, vote := range votes {
		if _, ok := elections[vote]; !ok {
			elections[vote] = 0
		}
		elections[vote]++
		if elections[vote] > len(votes)/2 {
			// Output 1 if the sequence contains an element
			// that appears strictly more than n/2 times
			return 1
		}
	}
	// And 0 otherwise.
	return 0
}

// Time Complexity: O(n)
func MajorityElementDivideAndConquer(votes []int) int {
	// To solve this problem you first split a given sequence
	// into halves and make two recursive calls.
	return getMajorityElement(votes)
}

func getMajorityElement(votes []int) int {
	votesLength := len(votes)

	// Base case we return single "major" elem
	if votesLength == 2 {
		if votes[0] == votes[0] {
			return votes[0]
		}

		return 0

	}
	if votesLength == 1 {
		return votes[0]
	}
	// DivideStep: we split votes slice
	// into two equals parts by calculating mid value
	mid := votesLength / 2
	votesLeft := votes[0:mid]
	votesRight := votes[mid:votesLength]

	// ConquerStep: we recursively calculate thr majority elem
	// on both halves and store them
	leftMajority := getMajorityElement(votesLeft)
	rightMajority := getMajorityElement(votesRight)

	if leftMajority == rightMajority {
		return leftMajority
	}

	leftFrequencyCount := countFrequency(votes, leftMajority)
	rightFrequencyCount := countFrequency(votes, rightMajority)

	if leftFrequencyCount > mid {
		return 1
		// return majorityInLower
	}

	if rightFrequencyCount > mid {
		return 1
		// return majorityInUpper
	}

	return 0

}

func countFrequency(seq []int, element int) int {
	count := 0
	for _, val := range seq {
		if element == val {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var a []int

	fmt.Scan(&n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(MajorityElement(a))
	// fmt.Println(MajorityElementDivideAndConquer(a))
}

/*
Input:
5
2 3 9 2 2
Output:
1

Input:
4
1 2 3 4
Output:
0

Input:
4
1 2 3 1
Output:
0
*/
