package main

import (
	"fmt"
)

func DanyaCalculator(number int) (int, []int) {
	// minNumOfOperations := make([]int, number+1)
	return 0, nil
}

// PrimitiveCalculator returns the minimum number of operations
// needed to obtain the number n starting from the number 1
func PrimitiveCalculator(number int) (int, []int) {
	/*
			Firstly we need to compute the minimum of operations;
			To do this we need to design a dynamic programming solution :/
			We can generalize the recurrences for this problem:
					  {	C(n/3) + 1 if n ⋮ 3;
			C(n) = min{ C(n/2) + 1 if n ⋮ 2;
					  {	C(n-1) + 1


			In other words min operations to reach(n)
			is min of min operations to reach(n-1) and
					  min operations to reach(n/2) and
		              min operations to reach(n/3)

			Example:
				n = 4,
				we can represent 4 = 4/2 * 2 or 4 = (4–1) + 1.
				C(4) = C(n/2) + 1 or C(4) = C(n-1) + 1.
			Because 4 is not divided by 3 so, we ignore C(n/3).
				n = 10
				numbers			0	1	2	3	4	5	6	7	8	9	10		// we can assume that it's our for loop
				operations		_	0	1	1	2	3	2		3	2	4|3		// minNumOfOperations
				memorization	_		1	1	2		2		4	3	5|9		// backtracking helper slice

			And to reconstruct the sequence we just simply go backwards
			of "memorization" slice printing the n (here 10) and get the value by in index n
			10 9 3 1 and reverse it :3
	*/
	minNumOfOperations := make([]int, number+1)
	// Fill minNumOfOperations of steps with max value
	for idx := range minNumOfOperations {
		minNumOfOperations[idx] = 9223372036854775807 // math.MaxInt
	}
	minNumOfOperations[0], minNumOfOperations[1] = 0, 0
	/*
		While we compute the minimum of operations,
		we need to backtracking to give the sequence of number.
		It contains all number of operations we need to reach from 0, 1, 2, … n.
	*/
	prev_num := make([]int, number+1)

	for i := 2; i < len(minNumOfOperations); i++ {
		// 3 * x
		if i%3 == 0 && minNumOfOperations[i] > minNumOfOperations[i/3]+1 {
			minNumOfOperations[i] = minNumOfOperations[i/3] + 1
			prev_num[i] = i / 3
		}
		// 2 * x
		// If it divisible by 2 and number of current operation greater (>) than
		if i%2 == 0 && minNumOfOperations[i] > minNumOfOperations[i/2]+1 {
			minNumOfOperations[i] = minNumOfOperations[i/2] + 1
			prev_num[i] = i / 2
		}
		// x + 1
		if minNumOfOperations[i] > minNumOfOperations[i-1]+1 {
			minNumOfOperations[i] = minNumOfOperations[i-1] + 1
			prev_num[i] = i - 1
		}
	}
	return minNumOfOperations[len(minNumOfOperations)-1], prev_num
}

func GreedyCalculator(number int) (int, []int) {
	minNumOperations := 0
	operations := []int{number}
	for number > 1 {
		if number%3 == 0 {
			number /= 3
			operations = append(operations, number)
		} else if number%2 == 0 {
			number /= 2
			operations = append(operations, number)
		} else {
			number--
			operations = append(operations, number)
		}
		minNumOperations++
	}

	Reverse(operations)

	return minNumOperations, operations
}

// Reverse is a helper function to reverse int slice cause idk
// what is wrong with sort.Sort(sort.Reverse(sort.IntSlice(operations)))
func Reverse(operations []int) {
	n := len(operations) - 1
	for i := 0; i < (n+1)/2; i++ {
		operations[i], operations[n-i] = operations[n-i], operations[i]
	}
}

func main() {
	var number int
	fmt.Scan(&number)

	// numberOp, operations := GreedyCalculator(number)
	// fmt.Println(numberOp)
	// for _, operation := range operations {
	// 	fmt.Printf("%d ", operation)
	// }

	numberOfOperations, numbers := PrimitiveCalculator(number)
	fmt.Println(numberOfOperations)

	answerDecrOrder := make([]int, 1)
	answerDecrOrder[0] = number
	for i := len(numbers) - 1; i > 1; {
		answerDecrOrder = append(answerDecrOrder, numbers[i])
		i = numbers[i]
	}
	Reverse(answerDecrOrder)
	for _, num := range answerDecrOrder {
		fmt.Printf("%v ", num)
	}
}

/*
Input:
1
Output:
0
1

Input:
5
Output:
3
1 2 4 5

Input:
96234
Output:
14
1 3 9 10 11 22 66 198 594 1782 5346 16038 16039 32078 96234
*/
