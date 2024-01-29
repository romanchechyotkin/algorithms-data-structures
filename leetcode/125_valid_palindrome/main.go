package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("roma"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	ss := []rune(s)
	
	left := 0
	right := len(s)-1

	for left < right {

		if !unicode.IsLetter(ss[left]) {
			left++
			continue
		}

		if !unicode.IsLetter(ss[right]) {
			right--
			continue
		}

		if strings.ToLower(string(ss[left])) != strings.ToLower(string(ss[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

