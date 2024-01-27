package main

import "fmt"

func main() {
	fmt.Println(binary_search([]int{1, 2, 3, 4, 8, 11, 23, 34, 45}, 1))
	fmt.Println(binary_search([]int{1, 2, 3, 4, 8, 11, 23, 34, 45}, 9))
	fmt.Println(binary_search([]int{1, 2, 3, 4, 8, 11, 23, 34, 45}, 45))
	fmt.Println(binary_search([]int{1, 2, 3, 4, 8, 11, 23, 34, 45}, 46))
}

func binary_search(nums []int, val int) bool {
	low := 0
	high := len(nums)-1
	var mid, guess int

	for low <= high {
		mid = low + (high-low)/2
		guess = nums[mid]
		
		if guess == val {
			return true
		} else if guess < val {
			low = mid + 1 
		} else {
			high = mid - 1 
		}
	}

	return false
}