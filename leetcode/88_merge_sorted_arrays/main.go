package main

import "log"

func main() {
	log.Println(merge([]int{1,2,3,0,0,0}, 3, []int{2, 5, 6}, 3))
}


// nums1.length == m + n
// nums2.length == n

// 1 2 3 0 0 0
//       |

// 2 5 6
// |


func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if len(nums1) != n + m {
		return nil
	}

	if len(nums2) != n {
		return nil
	}

	res := make([]int, 0, n+m)

	
	ptr1 := 0
	ptr2 := 0

	for ptr1 < m && ptr2 < n {
		val1 := nums1[ptr1]
		val2 := nums2[ptr2]

		if val1 < val2 {
			res = append(res, val1)
			ptr1++
		} else {
			res = append(res, val2)
			ptr2++
		}
	}

	if ptr1 == m {
		res = append(res, nums2[ptr2:]...)
	} else if ptr2 == n {
		res = append(res, nums1[ptr1:]...)
	}


	return res
}