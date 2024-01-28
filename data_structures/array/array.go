package main

import "fmt"

func main() {
	fmt.Println("arrays")

	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers[0], numbers[len(numbers)-1])
	fmt.Println(len(numbers), cap(numbers))

	numbers[1] = 0
	fmt.Println(numbers)
}