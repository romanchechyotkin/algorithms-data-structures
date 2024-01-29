package main

import "fmt"

type Node struct {
	val  any
	next *Node
}

func main() {
	header := NewLinkedList(1)
	fmt.Println(header)

	header.Append(2)
	header.Append(3)
	header.Append(4)

	header.Print()

	header.Pop()
	header.Print()

	header.Pop()
	header.Print()

	header.Pop()
	header.Print()
}

func NewLinkedList(val any) *Node {
	return &Node{
		val: val,
		next: nil,
	}
}

func (ll *Node) Append(val any) {
	tmp := ll

	for tmp.next != nil {
		tmp = tmp.next
	}
	
	tmp.next = &Node{
		val: val,
		next: nil,
	}
}

func (ll *Node) Pop() *Node {
	tmp := &Node{
		val: 0,
		next: ll,
	}
	prev := &Node{}
                                   
	for tmp.next != nil {
		prev = tmp
		tmp = tmp.next
	}

	prev.next = nil
	return prev
}

func (ll *Node) InsertAt(val any, idx int) {

}

func (ll *Node) Print() {
	tmp := ll

	for tmp != nil {
		fmt.Print(tmp.val, " ")
		tmp = tmp.next
	}

	fmt.Println()
}