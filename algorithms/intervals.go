package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(mergeIntervals([][]int{{2, 4}, {1, 5}, {8, 10}, {9, 14}, {15, 19}}))
	log.Println(mergeIntervals([][]int{{2, 8}, {1, 5}, {8, 10}, {9, 14}, {15, 19}})) 
	log.Println(mergeIntervals([][]int{{2, 8}, {1, 4}, {4, 5}, {7, 11}, {0, 1}})) 
}

// [ [1,5], [2,4], [8,10], [9,14], [15, 19] ]
func mergeIntervals(intervals [][]int) [][]int {
	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0]) // [ [1, 5],  [8, 10]]

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i] // [2, 4]
		lastItem := res[len(res)-1]
		
		if lastItem[1] >= interval[0] {
			lastItem[1] = max(lastItem[1], interval[1])
		} else {
			res = append(res, interval)
		}
	}


	return res
}
