package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key   int
	left  int
	right int
}

type Tree []Node

// SlowReadTree builds tree from the input
// Failed case #16/24: time limit exceeded (Time used: 2.19/1.50 ...
func (tr *Tree) SlowReadTree() *Tree {
	var n int
	fmt.Scan(&n)
	newTree := make(Tree, n)
	for i := 0; i < n; i++ {
		var key, left, right int
		fmt.Scan(&key, &left, &right)

		node := Node{
			key, left, right,
		}

		newTree[i] = node
	}
	return &newTree
}

// FastReadTree builds tree from the input
// Max time used: 0.38/1.50
func (tr *Tree) FastReadTree() *Tree {
	var n int
	fmt.Scan(&n)
	newTree := make(Tree, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan() && i < n; i++ {
		var key, left, right int
		line := scanner.Text()
		// strings.Split(s, " ")
		txt := strings.Split(line, " ")
		key, _ = strconv.Atoi(txt[0])
		left, _ = strconv.Atoi(txt[1])
		right, _ = strconv.Atoi(txt[2])

		node := Node{
			key, left, right,
		}
		newTree[i] = node
	}
	return &newTree
}

func (tr *Tree) RunTraversals() {
	tr.InorderTraversal(0)
	fmt.Println()
	tr.PreorderTraversal(0)
	fmt.Println()
	tr.PostorderTraversal(0)
}

// InorderTraversal prints the key of each Node in the in-order traversal of the tree.
func (tr *Tree) InorderTraversal(idx int) {
	if idx == -1 {
		return
	}

	tr.InorderTraversal((*tr)[idx].left)
	fmt.Printf("%d ", (*tr)[idx].key)
	tr.InorderTraversal((*tr)[idx].right)
}

// PreorderTraversal prints the key of each Node in the pre-order traversal of the tree.
func (tr *Tree) PreorderTraversal(idx int) {
	if idx == -1 {
		return
	}

	fmt.Printf("%d ", (*tr)[idx].key)
	tr.PreorderTraversal((*tr)[idx].left)
	tr.PreorderTraversal((*tr)[idx].right)
}

// PostorderTraversal prints the key of each Node in the post-order traversal of the tree.
func (tr *Tree) PostorderTraversal(idx int) {
	if idx == -1 {
		return
	}
	tr.PostorderTraversal((*tr)[idx].left)
	tr.PostorderTraversal((*tr)[idx].right)
	fmt.Printf("%d ", (*tr)[idx].key)
}

func main() {
	myTree := Tree{}

	myTree = *myTree.FastReadTree()

	myTree.RunTraversals()
}
