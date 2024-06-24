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
	_ = h
	a.Edges = []*Node[int]{b, c}
	b.Edges = []*Node[int]{d}
	c.Edges = []*Node[int]{d, e}
	d.Edges = []*Node[int]{e, g}
	e.Edges = []*Node[int]{f}
	f.Edges = []*Node[int]{g}

	fmt.Println(bfs(a, g))
}



func bfs[T comparable](start, end *Node[T]) bool {
	queue := make([]*Node[T], 0)
	queue = append(queue, start)

	for len(queue) != 0 {
		currNode := queue[0]
		fmt.Println(currNode.Value, " ")
		queue = queue[1:]

		if currNode.Value == end.Value {
			return true
		}

		queue = append(queue, currNode.Edges...)
	}

	return false
}
