package main

import "fmt"

// It was put in the thread Section
// https://www.youtube.com/watch?v=KxGRhd_iWuE

type Job struct {
	threadId       uint // id of the thread which processed this job
	startProcTime  uint // time, in seconds, when this job was processed
	processingTime uint // time, in seconds, necessary to process this job
}

type Thread struct {
	id   uint
	busy uint /* timeToFinishJob this thread is available,
	if busy == 0 or the time, in seconds,
	until the thread will be busy, if busy > 0 */
}

// rightChild() returns the index of the right child of index i.
func rightChild(idx uint) uint { return 2*idx + 1 }

// leftChild() returns the index of the left child of index i.
func leftChild(idx uint) uint { return 2*idx + 2 }

// processor() simulates the processing of jobs following the instructions of
// - the problem (see the corresponding programming assignment file),
// where nthd is the number of threads
func processor(arr *[]Job, nel, nthd uint) {
	pool := make([]Thread, nthd)

	for i := uint(0); i < nthd; i++ {
		pool[i].id = i
		pool[i].busy = 0
	}

	for i := uint(0); i < nel; i++ {
		(*arr)[i].threadId = pool[0].id
		(*arr)[i].startProcTime = pool[0].busy
		pool[0].busy += (*arr)[i].processingTime
		siftDown(&pool, nthd, 0)
	}

}

// siftDown() executes the necessary swap operations to garantee that the
// - job's array arr[i], ..., arr[nel-1] is a minHeap,
//  where nel is the number of elements in arr.
func siftDown(arr *[]Thread, nel, i uint) {
	j := i // index of the minimum value: arr[i], arr[lt] and arr[rt]

	lt := leftChild(i)
	if lt < nel {
		if (*arr)[lt].busy < (*arr)[j].busy || ((*arr)[lt].busy == (*arr)[j].busy && (*arr)[lt].id < (*arr)[j].id) {
			j = lt
		}
	}

	rt := rightChild(i)
	if rt < nel {
		if ((*arr)[rt].busy < (*arr)[j].busy) || (((*arr)[rt].busy == (*arr)[j].busy) && ((*arr)[rt].id < (*arr)[j].id)) {
			j = rt
		}
	}

	if i != j {
		// Swap
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		siftDown(arr, nel, j)
	}
}

func main() {
	// I really want to use parse input function,
	// but I somehow maneged to messed up
	var n, m uint
	fmt.Scan(&n, &m)

	arr := make([]Job, m)
	for i := range arr {
		fmt.Scan(&arr[i].processingTime)
	}

	processor(&arr, m, n)

	for i := range arr {
		fmt.Printf("%d %d\n", arr[i].threadId, arr[i].startProcTime)
	}
}
