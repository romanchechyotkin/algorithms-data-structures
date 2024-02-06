package main

import (
	"fmt"
)

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j')))
	fmt.Println(string(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z')))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	low := 0
	high := len(letters) - 1

	for low <= high {
		mid := low + (high-low)/2
		guess := letters[mid]
		if guess <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(letters) {
		return letters[0];
	}

	return letters[low]
}
