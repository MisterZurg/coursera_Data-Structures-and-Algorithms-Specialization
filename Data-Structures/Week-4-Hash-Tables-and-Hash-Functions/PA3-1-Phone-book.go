package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PhoneBook map[string]string

func (pb *PhoneBook) Add(number, name string) {
	(*pb)[number] = name
}

func (pb *PhoneBook) Del(number string) {
	delete(*pb, number)
}

func (pb *PhoneBook) Find(number string) string {
	if (*pb)[number] == "" {
		return "not found"
	}
	return (*pb)[number]
}

func ParseInput() [][]string {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	// The is slice of slice that contains order of commands
	// find 3839442
	// add 123456 me
	commands := make([][]string, n)
	for i := 0; scanner.Scan() && i < n; i++ {
		s := scanner.Text()
		commands[i] = strings.Split(s, " ")
	}
	return commands
}

func solution(pb PhoneBook, commands [][]string) {
	for _, cmd := range commands {
		switch cmd[0] {
		case "add":
			// cmd[1] is number, cmd[2] is name
			pb.Add(cmd[1], cmd[2])
		case "del":
			// cmd[1] is number
			pb.Del(cmd[1])
		case "find":
			// cmd[1] is number
			fmt.Println(pb.Find(cmd[1]))
		}
	}
}

func main() {
	commands := ParseInput()

	myPB := make(PhoneBook, len(commands))

	solution(myPB, commands)
}
