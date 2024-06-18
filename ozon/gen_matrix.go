package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// gen generates matrix NxM
func gen(n, m int) {
	rand.NewSource(time.Now().Unix())

	matrix := make([][]int, 0, n)
	counter := make(map[int]struct{})

	for i := 0; i < n; i++ {
		row := make([]int, 0, m)
		for j := 0; j < m; j++ {
			for {
				num := rand.Intn(100)
				_, ok := counter[num]

				if !ok {
					counter[num] = struct{}{}
					row = append(row, num)
					break
				}
	
			}
			
		}
		matrix = append(matrix, row)
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Print(matrix[row][col], " ")
		}
		fmt.Println()
	}

}

func main() {
	gen(3, 2)
	log.Println("------------------")
	gen(2, 3)
	log.Println("------------------")
	gen(4, 10)
}
