package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

// Insert At Beginning
func (l *LinkedList) InsertAtBeginning(val int) {
	newNode := &Node{Val: val}
	newNode.Next = l.Head
	l.Head = newNode
}

// Insert At End
func (l *LinkedList) InsertAtEnd(val int) {
	newNode := &Node{Val: val}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

// Insert After a Given Value
func (l *LinkedList) InsertAfter(target int, val int) {
	curr := l.Head
	for curr != nil && curr.Val != target {
		curr = curr.Next
	}
	if curr == nil {
		fmt.Println("Value not found")
		return
	}
	newNode := &Node{Val: val, Next: curr.Next}
	curr.Next = newNode
}

// Delete First Node
func (l *LinkedList) DeleteFirst() {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
}

// Delete Node By Value
func (l *LinkedList) DeleteValue(val int) {
	if l.Head == nil {
		return
	}

	if l.Head.Val == val {
		l.Head = l.Head.Next
		return
	}

	curr := l.Head
	for curr.Next != nil && curr.Next.Val != val {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

// Search Node
func (l *LinkedList) Search(val int) bool {
	curr := l.Head
	for curr != nil {
		if curr.Val == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

// Reverse Linked List
func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

// Print Linked List
func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}

	list.InsertAtBeginning(3)
	list.InsertAtBeginning(2)
	list.InsertAtBeginning(1)
	list.InsertAtEnd(4)
	list.InsertAtEnd(5)

	fmt.Print("Linked List: ")
	list.Print()

	fmt.Println("Search 4:", list.Search(4))

	list.DeleteValue(3)
	fmt.Print("After deleting 3: ")
	list.Print()

	list.InsertAfter(2, 10)
	fmt.Print("After inserting 10 after 2: ")
	list.Print()

	list.Reverse()
	fmt.Print("Reversed List: ")
	list.Print()
}
