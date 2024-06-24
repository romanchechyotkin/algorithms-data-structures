package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Edges []*Node[T]
}

//   / 2 \    / - - 7
// 1     /  4 \     |
//   \ 3   -   5 - 6     8

// stact  1

func (n *Node[T]) GetEdges() {
	fmt.Printf("from [%v]: ", n.Value)
	for _, e := range n.Edges {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {
	a := &Node[int]{Value: 1}
	b := &Node[int]{Value: 2}
	c := &Node[int]{Value: 3}
	d := &Node[int]{Value: 4}
	e := &Node[int]{Value: 5}
	f := &Node[int]{Value: 6}
	g := &Node[int]{Value: 7}
	h := &Node[int]{Value: 8}

	a.Edges = []*Node[int]{b, c}
	b.Edges = []*Node[int]{d}
	c.Edges = []*Node[int]{d, e}
	d.Edges = []*Node[int]{e, g}
	e.Edges = []*Node[int]{f}
	f.Edges = []*Node[int]{g}

	fmt.Println(dfs(a, h))
}

func dfs[T comparable](start, end *Node[T]) bool {
	stack := make([]*Node[T], 0)
	stack = append(stack, start)

	for len(stack) != 0 {
		currNode := stack[len(stack)-1]
		fmt.Println(currNode.Value, " ")
		stack = stack[:len(stack)-1]

		if currNode.Value == end.Value {
			return true
		}

		stack = append(stack, currNode.Edges...)
	}

	return false
}
