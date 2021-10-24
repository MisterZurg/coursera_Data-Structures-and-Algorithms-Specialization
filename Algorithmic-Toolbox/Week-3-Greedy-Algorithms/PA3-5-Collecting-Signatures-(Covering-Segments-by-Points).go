package main

import "fmt"

func CollectingSignatures(segments [][]int) (int, []int) {
	// Sort segments slice according to the end points
	sortByLastElement(segments)

	// Create points slice to store the common points in the segments
	// set the point to the first end point i-e shortest end point
	points := make([]int, 1)
	points[0] = segments[0][1]
	pPointer := 0
	// if the point is not in the segment
	// update the point to the end point of the current segment
	// store it in the vector
	for i := 1; i < len(segments); i++ {
		if points[pPointer] < segments[i][0] {
			points = append(points, segments[i][1])
			pPointer++
		}
	}
	return len(points), points
}

func main() {
	var segmentsNumber int
	fmt.Scan(&segmentsNumber)

	segments := fillSegments(segmentsNumber)

	num, points := CollectingSignatures(segments)
	fmt.Println(num)
	for _, point := range points {
		fmt.Printf("%d ", point)
	}
}

func fillSegments(segmentsNumber int) [][]int {
	segments := make([][]int, segmentsNumber)
	// Initializing nested slice
	for i := range segments {
		segments[i] = make([]int, 2)
	}

	// Filling nested slice with values
	for i := range segments {
		for j := range segments[i] {
			fmt.Scan(&segments[i][j])
		}
	}
	return segments
}

func sortByLastElement(unsortedSegments [][]int) {
	// Very stupid sort
	for i := 0; i < len(unsortedSegments)-1; i++ {
		for j := 0; j < len(unsortedSegments)-1; j++ {
			if unsortedSegments[j][1] > unsortedSegments[j+1][1] {
				unsortedSegments[j], unsortedSegments[j+1] = unsortedSegments[j+1], unsortedSegments[j]
			}
		}
	}
}

/*

Test case #3/15:
Input:
100
41 42
52 52
63 63
80 82
78 79
35 35
22 23
31 32
44 45
81 82
36 38
10 12
1 1
23 23
32 33
87 88
55 56
69 71
89 91
93 93
38 40
33 34
14 16
57 59
70 72
36 36
29 29
73 74
66 68
36 38
1 3
49 50
68 70
26 28
30 30
1 2
64 65
57 58
58 58
51 53
41 41
17 18
45 46
4 4
0 1
65 67
92 93
84 85
75 77
39 41
15 15
29 31
83 84
12 14
91 93
83 84
81 81
3 4
66 67
8 8
17 19
86 87
44 44
34 34
74 74
94 95
79 81
29 29
60 61
58 59
62 62
54 56
58 58
79 79
89 91
40 42
2 4
12 14
5 5
28 28
35 36
7 8
82 84
49 51
2 4
57 59
25 27
52 53
48 49
9 9
10 10
78 78
26 26
83 84
22 24
86 87
52 54
49 51
63 64
54 54

Correct output:
43
1 4 5 8 9 10 14 15 18 23 26 28 29 30 32 34 35 36 40 41 44 46 49 52 54 56 58 61 62 63 65 67 70 74 77 78 79 81 84 87 91 93 95
 (Time used: 0.00/1.50, memory used: 11530240/2147483648.)
*/
