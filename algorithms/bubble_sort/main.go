package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{21, 3, 4, 5, 0, 1, -1, 22}))
	fmt.Println(bubbleSort([]int{9, 8, 7, 6, 5, 4, 2, 1}))
	fmt.Println(bubbleSort([]int{1, 2, 3, 4, 5, 6, 0}))
} 

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-(i+1); j++ {
			left := nums[j]
			right := nums[j+1]
			
			if left > right {
				nums[j] = right
				nums[j+1] = left
			}
		}
	}	


	return nums
}	