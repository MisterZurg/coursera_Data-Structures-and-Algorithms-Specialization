package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MIN_VALUE = -2147483648
	MAX_VALUE = 2147483647
)

type node struct {
	key   int
	left  int
	right int
}

type tree []node

// ReadTree builds tree from the input
func (tr *tree) fastReadTree() *tree {
	var n int
	fmt.Scan(&n)
	newTree := make(tree, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan() && i < n; i++ {
		var key, left, right int
		line := scanner.Text()

		txt := strings.Split(line, " ")
		key, _ = strconv.Atoi(txt[0])
		left, _ = strconv.Atoi(txt[1])
		right, _ = strconv.Atoi(txt[2])

		node := node{
			key, left, right,
		}
		newTree[i] = node
	}
	return &newTree
}

func (tr tree) check(idx, min, max int) bool {
	if idx == -1 {
		return true
	}

	if tr[idx].key <= min || tr[idx].key >= max {
		return false
	}
	return tr.check(tr[idx].left, min, tr[idx].key) && tr.check(tr[idx].right, tr[idx].key, max)
}

func (tr tree) isBinarySearchTree() bool {
	if len(tr) == 0 {
		return true
	}

	return tr.check(0, MIN_VALUE, MAX_VALUE)
}

func (tr *tree) Solution() {
	if tr.isBinarySearchTree() {
		fmt.Println("CORRECT")
		return
	}
	fmt.Println("INCORRECT")
	return
}

func main() {
	myTree := tree{}
	myTree = *myTree.fastReadTree()

	myTree.Solution()
}

/*
Input:
3
2 1 2
1 -1 -1
3 -1 -1
Output:
CORRECT

Sample 2.
Input:
3
1 1 2
2 -1 -1
3 -1 -1
Output:
INCORRECT
*/
