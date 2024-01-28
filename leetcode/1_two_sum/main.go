package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 9}, 9))
	fmt.Println(twoSum([]int{0, 4, 2}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, n := range nums {
		
		if v, ok := hashMap[n]; !ok {
			hashMap[target-n] = i
		} else {
			return []int{v, i}
		}

	}

	return nil
}