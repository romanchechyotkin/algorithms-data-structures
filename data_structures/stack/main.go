package main

import (
	"fmt"
)

type Node struct {
	val  any
	prev *Node
}

type Stack struct {
	head *Node
	l    int
}

func main() {
	stack := NewStack()
	fmt.Println(stack.l)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.l)

	stack = stack.Pop()
	fmt.Println(stack.head)
	stack = stack.Pop()
	fmt.Println(stack.head)
	stack = stack.Pop()
	fmt.Println(stack)
}

func NewStack() *Stack {
	return &Stack{
		head: nil,
		l:    0,
	}
}

func (s *Stack) Push(val any) {

	n := &Node{
		val: val,
	}

	if s.head == nil {
		s.head = n
		s.l++
		return
	}

	n.prev = s.head
	s.head = n
	s.l++
}

func (s *Stack) Pop() *Stack {
	node := s.head 
	s.head = node.prev 
	s.l--

	return s
}
