package main

import "fmt"

type Table struct {
	parent    uint
	rowNumber uint
}

type Tables []Table

// MakeTable creates the i-th table in tbl and registers the number of rows in it.
func (tbls *Tables) MakeTable(i, rowNumber uint) {
	(*tbls)[i].parent = i
	(*tbls)[i].rowNumber = rowNumber
}

// Find finds the destination of the i-th table in tbl and compresses the path between them.
func (tbls *Tables) Find(i uint) uint {
	j := (*tbls)[i].parent
	if j != i {
		j = tbls.Find(j)
	}
	return j
}

// Merge if the destination of i is different of the destination of j in tbl,
// copies all the rows from the first table to the second table, then clears the first table and,
// instead of real data, puts a symbolic link to the second table into it.
// Returns the maximum of the sizes of all tables in tbl after the merge operation,
// where pmax is the maximum value before this operation.
func (tbls *Tables) Merge(i, j, pmax uint) uint {
	i_ID := tbls.Find(i)
	j_ID := tbls.Find(j)

	if i_ID == j_ID {
		return pmax
	}

	(*tbls)[j_ID].parent = i_ID

	(*tbls)[i_ID].rowNumber += (*tbls)[j_ID].rowNumber
	temp := (*tbls)[i_ID].rowNumber
	(*tbls)[j_ID].rowNumber = 0

	if temp > pmax {
		return temp
	}
	return pmax
}

func main() {
	// n tables stored
	// m the number of merge queries to perform
	var n, m uint
	fmt.Scan(&n, &m)

	tables := make(Tables, n)
	pmax := uint(0)

	var rowNum uint

	for i := uint(0); i < n; i++ {
		fmt.Scan(&rowNum)
		tables.MakeTable(i, rowNum)

		if rowNum > pmax {
			pmax = rowNum
		}
	}

	var d, s uint
	for i := uint(0); i < m; i++ {
		fmt.Scan(&d, &s)
		pmax = tables.Merge(d-1, s-1, pmax)

		fmt.Printf("%d\n", pmax)
	}
}

/*
TC 1
5 5
1 1 1 1 1
3 5
2 4
1 4
5 4
5 3


TC 2
6 4
10 0 5 0 3 3
6 6
6 5
5 4
4 3
*/
