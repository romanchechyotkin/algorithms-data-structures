package main

import "fmt"


func main()  {
	fmt.Println(validAnagram("cat", "rat"))	
	fmt.Println(validAnagram("anagram", "nagaram"))	
	fmt.Println(validAnagram("aacc", "ccac"))	
}

func validAnagram(s, t string) bool {
	hashMap := make(map[rune]int)

	for _, v := range s {
		hashMap[v]++
	}

	for _, v := range t {
		hashMap[v]--
	}

	for _, v := range hashMap {
		if v != 0 {
			return false
		}
	}
	
	return true
}

