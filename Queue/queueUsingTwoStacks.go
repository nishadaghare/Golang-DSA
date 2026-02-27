package main

import "fmt"

type Queue struct {
	inStack  []int
	outStack []int
}

// Enqueue operation
func (q *Queue) Enqueue(x int) {
	q.inStack = append(q.inStack, x)
}

// Dequeue operation
func (q *Queue) Dequeue() int {
	if len(q.outStack) == 0 {
		// Transfer elements
		for len(q.inStack) > 0 {
			n := len(q.inStack) - 1
			q.outStack = append(q.outStack, q.inStack[n])
			q.inStack = q.inStack[:n]
		}
	}

	if len(q.outStack) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}

	n := len(q.outStack) - 1
	val := q.outStack[n]
	q.outStack = q.outStack[:n]
	return val
}

func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println(q.Dequeue()) // 10
	fmt.Println(q.Dequeue()) // 20

	q.Enqueue(40)

	fmt.Println(q.Dequeue()) // 30
	fmt.Println(q.Dequeue()) // 40
}
