package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoCrystalBalls([]bool{false, false, true, false, false, true}))
	fmt.Println(twoCrystalBalls([]bool{false, false, false, false, false, true, true, true, true}))
	fmt.Println(twoCrystalBalls([]bool{false, false, true}))
	fmt.Println(twoCrystalBalls([]bool{true, true, true}))
	fmt.Println(twoCrystalBalls([]bool{false, false, true, true}))
	fmt.Println(twoCrystalBalls([]bool{false, false, false, false, false, false, false, true}))
}

func twoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(breaks))))
	
	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for j := i; j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}
