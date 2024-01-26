package main

import "fmt"

func main() {
	fmt.Println(search([]int{1, 2, 3}, 1))
	fmt.Println(search([]int{1, 2, 3}, 4))
}

func search(nums []int, val int) bool {
	for num := range nums {
		if num == val {
			return true
		}
	}

	return false
}