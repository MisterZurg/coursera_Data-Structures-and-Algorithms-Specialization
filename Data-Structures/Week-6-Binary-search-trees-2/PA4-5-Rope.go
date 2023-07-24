package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rope string

func (rp *Rope) Process(i, j, k int) {
	substring := (*rp)[i : j+1]
	*rp = (*rp)[:i] + (*rp)[j+1:]

	if k == 0 {
		*rp = substring + *rp
	} else {
		*rp = (*rp)[:k] + substring + (*rp)[k:]
	}
}

func ParseInput() (Rope, int) {
	/*
		The first line contains the initial string 𝑆.
		The second line contains the number of queries 𝑞.
		Next 𝑞 lines contain triples of integers 𝑖, 𝑗, 𝑘.
	*/
	var initialString Rope
	var queriesNumber int

	fmt.Scan(&initialString, &queriesNumber)

	return initialString, queriesNumber
}

func (rp *Rope) getOperations() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	q, _ := strconv.Atoi(scanner.Text())

	for idx := 0; scanner.Scan() && idx < q; idx++ {
		line := scanner.Text()
		txt := strings.Split(line, " ")
		i, _ := strconv.Atoi(txt[0])
		j, _ := strconv.Atoi(txt[1])
		k, _ := strconv.Atoi(txt[2])
		rp.Process(i, j, k)
	}

}

func main() {
	var initialString Rope
	fmt.Scan(&initialString)
	initialString.getOperations()

	fmt.Print(initialString)
}
