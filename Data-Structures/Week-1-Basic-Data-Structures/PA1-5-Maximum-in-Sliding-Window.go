package main

import (
	"errors"
	"fmt"
)

type DequeNode struct {
	value int
	left  *DequeNode
	right *DequeNode
}

type Deque struct {
	Capacity                int
	CurrentNumberOfElements int
	Front                   *DequeNode
	Back                    *DequeNode
}

var EmptyDeque = errors.New("DequeIsEmpty")

func (q *Deque) PushFront(elem int) {
	// Если это превый добавляемый элемент
	// 1 -> []
	if q.CurrentNumberOfElements == 0 {
		q.CurrentNumberOfElements++

		newDqNode := DequeNode{
			value: elem,
			left:  nil,
			right: nil,
		}

		q.Front = &newDqNode
		q.Back = &newDqNode
		// Если в очереди уже что-то есть и места достаточно
		// 3 - > [2 <|> 1]
	} else if q.Capacity > q.CurrentNumberOfElements {
		q.CurrentNumberOfElements++
		newDqNode := DequeNode{
			value: elem,
			left:  nil,
			right: q.Front,
		}
		// Обновляем прошлый первый элемент
		q.Front.left = &newDqNode
		q.Front = &newDqNode
	}
}

func (q *Deque) PushBack(elem int) {
	// Если это превый добавляемый элемент
	// 1 -> []
	if q.CurrentNumberOfElements == 0 {
		q.CurrentNumberOfElements++

		newDqNode := DequeNode{
			value: elem,
			left:  nil,
			right: nil,
		}

		q.Front = &newDqNode
		q.Back = &newDqNode
		// Если в очереди уже что-то есть и места достаточно
		// [2 <|> 1] <- 3
	} else if q.Capacity > q.CurrentNumberOfElements {
		q.CurrentNumberOfElements++
		newDqNode := DequeNode{
			value: elem,
			left:  q.Back,
			right: nil,
		}
		// Обновляем прошлый послежний элемент
		q.Back.right = &newDqNode
		q.Back = &newDqNode
	}
}

func (q *Deque) PopFront() *DequeNode {
	if q.CurrentNumberOfElements > 1 {
		front := q.Front
		q.Front = front.right

		q.Front.left = nil
		front.right = nil

		q.CurrentNumberOfElements--
		return front
		// Если элемент последний
	} else if q.CurrentNumberOfElements == 1 {
		lastElem := q.Front
		q.Front = nil
		q.Back = nil

		q.CurrentNumberOfElements--

		return lastElem
	}
	return nil
}

func (q *Deque) PopBack() *DequeNode {
	if q.CurrentNumberOfElements > 1 {
		back := q.Back
		q.Back = back.left

		q.Back.right = nil
		back.left = nil

		q.CurrentNumberOfElements--
		return back
		// Если элемент последний
	} else if q.CurrentNumberOfElements == 1 {
		lastElem := q.Back
		q.Front = nil
		q.Back = nil

		q.CurrentNumberOfElements--

		return lastElem
	}
	return nil
}

func (q *Deque) LookFront() (int, error) {
	if q.IsEpty() {
		return 0, EmptyDeque
	}
	return q.Front.value, nil
}

func (q *Deque) LookBack() (int, error) {
	if q.IsEpty() {
		return 0, EmptyDeque
	}
	return q.Back.value, nil
}

func (q *Deque) IsEpty() bool {
	return q.CurrentNumberOfElements > 0
}

func ParseInput5() ([]int, int) {
	var n int
	fmt.Scan(&n)

	notAwindow := make([]int, n)
	for i := range notAwindow {
		fmt.Scan(&notAwindow[i])
	}

	var window int
	fmt.Scan(&window)
	return notAwindow, window
}

func main() {
	sequence, window := ParseInput5()
	//answer := FindMaxInSequence(sequence, window)
	//for _, elem := range answer {
	//	fmt.Printf("%d ", elem)
	//}

	FindMaxInSequence(sequence, window)

	//myDq := Deque{Capacity: 3}
	//myDq.PushBack(3)
	//myDq.PushBack(2)
	//myDq.PushBack(1)
	//
	//fmt.Println(myDq.PopFront())
	//fmt.Println(myDq.LookFront())
	//
	//fmt.Println(myDq.PopFront())
	//fmt.Println(myDq.LookFront())
	//fmt.Println(myDq.PopFront())
	//fmt.Println(myDq.PopFront())
	//fmt.Println(myDq.LookFront())
}

// FindMaxInSequence решает задачу, поиска в максимума в подмассиве, используя Deque
// Algorithm:
// - Initialise the first K elements of the array into the deque.
// - Iterate over the input array and for each step:
// - Consider only the indices of the elements in the current window.
// - Pop-out all the indices of elements smaller than the current element, since their value will be less than the current element.
// - Push the current element into the deque.
// - Push the first element of the deque i.e. deque[0] into the output array.
func FindMaxInSequence(sequence []int, window int) {
	myDq := Deque{Capacity: window}
	// NumberOfWindows := len(sequence) - window + 1
	// answer := make([]int, NumberOfWindows)
	// ansIdx := 0

	// Пушим первый элемент в очередь

	// После этого, каждый раз добавляя новый элемент нам нужно проверить
	// if sequence[i] >= sequence[Dequeue.LookBack()] {
	//     Dequeue.PopBack
	// }

	// После того как мы начинаем двигать окно, то бишь каждую следующую итерацию цикла
	// Мы выводим максимум, который находится вначале очереди
	// fmt.Println(Dequeue.LookАкщте())
	// Мы должны удалять элемент, из нёё индекс которого уже не в окне index <= i - k
	//

	// Заполним Дек, на размер окна
	for i := 0; i < window; i++ {
		// Гарантированно пихаем первый
		if i == 0 {
			myDq.PushBack(i)
			// После этого, каждый раз добавляя новый элемент нам нужно проверить
		} else {
			dqBack, err := myDq.LookBack()
			// В случае пустой очереди
			if err != nil {
				continue
			}
			if sequence[i] >= sequence[dqBack] {
				// fmt.Println("Условие ", sequence[myDq.LookBack()] < sequence[i])
				for !myDq.IsEpty() && sequence[i] >= sequence[dqBack] {
					myDq.PopBack()
					dqBack, err = myDq.LookBack()
					// В случае пустой очереди
					if err != nil {
						continue
					}
				}
				// Вставляем индекс текущего элемента
				myDq.PushBack(i)
			}
		}
		// fmt.Println("index=", i)
		// fmt.Println("max",myDq.LookFront())
	}
	// Тут мы уже двигаем окно
	for i := window; i < len(sequence); i++ {
		// Выводим максимум, прошлого окна
		// элемент из последовательности по индексу вначале очереди
		fmt.Println(sequence[myDq.Front.value])
		// fmt.Println("index=", i)
		// fmt.Println(i, myDq.LookFront() - window)
		// Мы должны удалять элемент, из нёё индекс которого уже не в окне index <= i - k
		dqFront, err := myDq.LookFront()
		if err != nil {
			continue
		}
		if dqFront <= i-window {
			myDq.PopFront()
		}
		dqBack, err := myDq.LookBack()
		if sequence[dqBack] < sequence[i] {
			for !myDq.IsEpty() && sequence[dqBack] <= sequence[i] {
				myDq.PopBack()

				dqBack, err = myDq.LookBack()
				// В случае пустой очереди
				if err != nil {
					continue
				}
			}
			// Вставляем индекс текущего элемента
			myDq.PushBack(i)
		}
	}
	// return answer
}

/*
6
20 5 3 8 6 15
3


6
1 3 4 8 6 15
3
*/
