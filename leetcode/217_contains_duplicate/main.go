package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate( []int{1, 2, 2, 4, 5} ))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 2, 3, 4, 1}))
}

func containsDuplicate(nums []int) bool {
    m := make(map[int]int)

    for _, key := range nums {
        if m[key] == 1 {
            return true
        } else {
            m[key] = 1                
        }
    } 

    return false

}