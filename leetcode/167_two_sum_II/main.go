package main

import "fmt"

func main()  {
	fmt.Println(twoSum([]int{1, 2, 3}, 3))
	fmt.Println(twoSum([]int{0, 2, 3}, 3))
	fmt.Println(twoSum([]int{-1, 0, 1}, 1))
	fmt.Println(twoSum([]int{0, 1}, 1))
}

func twoSum(nums []int, target int) []int {
 	left := 0
 	right := len(nums) - 1 

	for {
		if nums[left] + nums[right] < target {
			left += 1
		} else if nums[left] + nums[right] > target {
			right -= 1
		} else {
			return []int{left+1, right+1}
		}
	}
}