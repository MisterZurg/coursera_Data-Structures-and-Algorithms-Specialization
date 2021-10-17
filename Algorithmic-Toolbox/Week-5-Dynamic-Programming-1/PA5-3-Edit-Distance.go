package main

import (
	"fmt"
)

func EditDistance(wordOne, wordTwo string) (int, [][]int) {
	// It first fills in the first column and the first row of the dynamic programming matrix
	// and then it continues filling it up by computing the cost of moving to vertex (i, j)
	// using insertion, deletion, or mismatch or match or in other words,
	// exploring all possibility. Moving to the vertex i, j using vertical edge,
	// horizontal edge, and diagonal edge.
	//
	// And then it finds out which of these possibilities results in the minimum edit distance.

	// table will store results of sub-problems
	table := make([][]int, len(wordOne)+1)
	for i := range table {
		table[i] = make([]int, len(wordTwo)+1)
	}
	// Print table
	// printTable(table)

	// Filling distances table(i, 0) in the first column of this matrix.
	// It is easy because indeed we are comparing an i-prefix of string A against a 0-prefix
	// of string D and therefore this edit distance for i-prefix will be equal to i.
	for i := range table {
		table[i][0] = i
	}
	// Print table after filling first column
	// printTable(table)
	// Similarly we can easily fill the first row in this matrix.
	for j := range table[0] {
		table[0][j] = j
	}
	// Print table after filling first row
	// printTable(table)

	// Filling the matrix
	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[0]); j++ {
			// Same logic from lecture
			insertion := table[i][j-1] + 1
			deletion := table[i-1][j] + 1
			match := table[i-1][j-1]
			mismatch := table[i-1][j-1] + 1

			if wordOne[i-1] == wordTwo[j-1] {
				table[i][j] = Min(insertion, Min(deletion, match))
			} else if wordOne[i-1] != wordTwo[j-1] {
				table[i][j] = Min(insertion, Min(deletion, mismatch))
			}
		}
	}
	// Print table with filled distances
	// printTable(table)

	return table[len(wordOne)][len(wordTwo)], table
}

func OutputAligment(i, j int) {
	if i == 0 && j == 0 {
		return
	}

}

// Min is a helper function to return min number of two ints
func Min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

// printTable is a helper function to output the table
func printTable(table [][]int) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	for i := range table {
		fmt.Println(table[i])
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
}

func main() {
	var wordOne, wordTwo string
	fmt.Scan(&wordOne, &wordTwo)

	editDistance, _ := EditDistance(wordOne, wordTwo)
	fmt.Println(editDistance)
}

/*
Input:
ab
ab
Output:
0

Input:
short
ports

Output:
3

Input:
editing
distance

Output:
5
*/
