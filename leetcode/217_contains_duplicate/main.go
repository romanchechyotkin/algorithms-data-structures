package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 2, 4, 5}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 2, 3, 4, 1}))
}

func containsDuplicate(nums []int) bool {
	counter := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := counter[num]; ok {
			return true
		} else {
			counter[num] = struct{}{}
		}
	}

	return false
}
