package main

import "fmt"

type Node struct {
	val  any
	next *Node
}

func main() {
	head := NewQueue(1)
	head.Print()

	head.Push(2)
	head.Push(3)
	head.Push(4)
	head.Print()

	head = head.Pop()
	head.Print()
}

func NewQueue(val any) *Node {
	return &Node{
		val:  val,
		next: nil,
	}
}

func (n *Node) Pop() *Node {
	h := n 
	n = h.next
	h.next = nil  
	return n
}

func (n *Node) Push(val any) {
	tmp := n

	for tmp.next != nil {
		tmp = tmp.next 
	}

	tmp.next = &Node{
		val: val,
	}
}

func (n *Node) Print() {
	tmp := n

	for tmp != nil {
		fmt.Print(tmp.val, " ")
		tmp = tmp.next 
	}

	fmt.Println()
}