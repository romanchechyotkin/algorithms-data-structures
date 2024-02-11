package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	row int
	col int
}

var dir = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func main() {
	f, err := os.Open("algorithms/maze_solver/maze.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	maze := make([][]string, 0)

	for s.Scan() {
		line := s.Text()
		col := strings.Split(line, "")

		maze = append(maze, col)
	}

	for _, row := range maze {
		fmt.Println(row)
	}

	moves := mazeSolver(maze, "#", Point{
		row: 2,
		col: 1,
	}, Point{
		row: 0,
		col: 5,
	})

	for _, v := range moves {
		fmt.Println(v)
	}
}

func mazeSolver(maze [][]string, wall string, start Point, finish Point) []Point {
	path := make([]Point, 0)
	seen := make([][]bool, len(maze))

	for i, _ := range seen {
		defaultRaw := make([]bool, len(maze[0]))
		seen[i] = defaultRaw
	} 

	walk(maze, wall, start, finish, seen, &path) 
		
	return path
}

func walk(maze [][]string, wall string, curr Point, finish Point, seen [][]bool, path *[]Point) bool {
	// base case
	// 1. its wall #
	// 2. out of bounds
	// 3. the finish Point
	// 4. if cell is seen

	if curr.row >= len(maze) || curr.row < 0 || curr.col >= len(maze[0]) || curr.col < 0 {
		return false
	}

	if maze[curr.row][curr.col] == wall {
		return false
	}

	if curr == finish {
		*path = append(*path, finish)
		return true
	}

	if seen[curr.row][curr.col] {
		return false
	}

	// pre
	seen[curr.row][curr.col] = true
	*path = append(*path, curr)

	// recursive case
	for i := 0; i < len(dir); i++ {
		if walk(maze, wall, Point{
			row: curr.row + dir[i][0],
			col: curr.col + dir[i][1],
		}, finish, seen, path) {
			return true
		}
	}

	// post
	slice := *path
	*path = slice[:len(slice)-1]

	return false
}
