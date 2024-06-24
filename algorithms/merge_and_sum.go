package main

import "fmt"

func main() {
	fmt.Println(mergeAndSum([]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6})) // 5 7 9 4 5 6
}

// 1 2 3 4 5 6
// |

// 4 5 6
//     |

// 5 7 9 4 5 6

func mergeAndSum(nums1, nums2 []int) []int {
	ptr1 := 0
	ptr2 := 0

	res := make([]int, 0, len(nums1)+len(nums2))

	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		val1 := nums1[ptr1]
		val2 := nums2[ptr2]
		sumVal := val1 + val2

		res = append(res, sumVal)

		ptr1++
		ptr2++
	}

	if ptr1 != len(nums1) {
		res = append(res, nums2...)
	} else if ptr2 != len(nums2) {
		res = append(res, nums1...)
	}
 
	return res
}
