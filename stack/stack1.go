package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

// Push element to stack
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// Pop element from stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	element := s.items[index]
	s.items = s.items[:index]

	return element, nil
}

// Peek top element
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Check if empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size of stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Size:", stack.Size())

	val, _ := stack.Pop()
	fmt.Println("Popped:", val)

	top, _ := stack.Peek()
	fmt.Println("Top:", top)
}