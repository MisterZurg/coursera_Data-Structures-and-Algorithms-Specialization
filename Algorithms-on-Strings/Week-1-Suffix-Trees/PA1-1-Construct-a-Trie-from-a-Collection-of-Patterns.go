package main

import (
	"fmt"
)

type Patterns []string

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	// TrieNode is an alphabet letter so we can replace []rune
	number   uint // output helper
	children [26]*TrieNode
	end      bool
}

func TrieConstruction(patterns Patterns) Trie {
	trie := Trie{root: &TrieNode{}}
	number := uint(1)
	for _, pattern := range patterns {
		var currentNode *TrieNode
		currentNode = trie.root

		for _, currentSymbol := range pattern {
			//fmt.Println(currentSymbol)
			if currentNode.children[currentSymbol-65] != nil {
				currentNode = currentNode.children[currentSymbol-65]
			} else {
				newNode := &TrieNode{number: number}
				number++
				currentNode.children[currentSymbol-65] = newNode
				currentNode = newNode
			}
		}
	}

	return trie
}

func (tr *Trie) PreOrderTraversal() {
	q := []*TrieNode{tr.root}

	for len(q) > 0 {
		currentNode := q[len(q)-1]
		q = q[:len(q)-1]
		for i, ch := range currentNode.children {
			if ch != nil {

				fmt.Printf("%d->%d:%c\n", currentNode.number, ch.number, getLetterFromIndex(i))
				q = append(q, ch)
			}
		}
	}
}

func getLetterFromIndex(i int) rune {
	return rune(i) + 65
}

func main() {
	var n int
	fmt.Scan(&n)

	patterns := make(Patterns, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&patterns[i])
	}
	trie := TrieConstruction(patterns)

	trie.PreOrderTraversal()

	//fmt.Println('a')
	//fmt.Println('z')
	//fmt.Println('A')
	//fmt.Println('Z')
}

/*

1
ATA


*/
