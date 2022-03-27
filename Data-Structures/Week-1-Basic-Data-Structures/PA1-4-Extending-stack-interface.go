package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyStack struct {
	currentMax int
	elements   []int
}

func main() {
	var numberOfCommands int
	mySt := MyStack{}

	fmt.Scan(&numberOfCommands)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < numberOfCommands; i++ {
		scanner.Scan()
		s := scanner.Text()
		command := strings.Split(s, " ")

		switch command[0] {
		case "push":
			value, _ := strconv.Atoi(command[1])
			mySt.Push(value)
		case "pop":
			mySt.Pop()
			mySt.UpdateCurrentMax()
		case "max":
			fmt.Println(mySt.Max())
		}
	}

}

func (ms *MyStack) Push(value int) {
	if ms.currentMax < value {
		ms.currentMax = value
	}
	ms.elements = append(ms.elements, value)
}

func (ms *MyStack) Pop() {
	ms.elements = ms.elements[:len(ms.elements)-1]
}

func (ms *MyStack) Max() int {
	return ms.currentMax
}

func (ms *MyStack) UpdateCurrentMax() {
	ms.currentMax = ms.elements[0]

	for _, elem := range ms.elements {
		if ms.currentMax < elem {
			ms.currentMax = elem
		}
	}
}

/*
Sample 1.
Input:
5
push 2
push 1
max
pop
max
Output:
2
2

Sample 2.
Input:
5
push 1
push 2
max
pop
max
Output:
2
1

Sample 3.
Input:
10
push 2
push 3
push 9
push 7
push 2
max
max
max
pop
max
Output:
9
9
9
9

Sample 4.
Input:
3
push 1
push 7
pop
Output:

Sample 5.
Input:
6
push 7
push 1
push 7
max
pop
max
Output:
7
7

*/
