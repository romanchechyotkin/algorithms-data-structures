package main

import "log"

func main() {
	log.Println(mergeArrays([]int{1, 4, 6}, []int{2, 3, 5}))
	log.Println(mergeArrays([]int{2, 3, 5}, []int{1}))
	log.Println(mergeArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	log.Println(mergeArrays([]int{}, []int{4, 5, 6}))
	log.Println(mergeArrays([]int{1, 2, 3}, []int{1, 2, 3}))
	log.Println(mergeArrays([]int{1, 3, 4, 7, 8}, []int{1, 3, 4, 5, 6, 7}))
}

// 1 2 3 4 5 6

// 1 2 3
// |

// 4 5 6
// |

// O(n)
func mergeArrays(arr1, arr2 []int) []int {
	l := len(arr1) + len(arr2)
	res := make([]int, 0, l)

	ptr1 := 0
	ptr2 := 0

	for ptr1 < len(arr1) && ptr2 < len(arr2) {
		val1 := arr1[ptr1]
		val2 := arr2[ptr2]

		if val1 > val2 {
			res = append(res, val2)
			ptr2++
		} else {
			res = append(res, val1)
			ptr1++
		}

	}

	if ptr1 < len(arr1) {
		res = append(res, arr1[ptr1:]...)
	} else if ptr2 < len(arr2) {
		res = append(res, arr2[ptr2:]...)
	}

	return res
}
