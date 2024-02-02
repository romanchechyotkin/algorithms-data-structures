package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{1, 2, 3}))
	fmt.Println(missingNumber([]int{3, 2, 1}))
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)

	var sum int
	for i := 0; i < n+1; i++ {
		sum += i
	}	

	for _, v := range nums {
		sum -= v
	}

	return sum
}