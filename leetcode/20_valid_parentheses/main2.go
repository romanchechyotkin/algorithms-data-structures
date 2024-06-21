package main

import "log"

func main() {
	log.Println(valid_parenteses("{{([])}}")) // true
	log.Println(valid_parenteses("{{([}}"))   // false
	log.Println(valid_parenteses("{{])}}"))   // false
	log.Println(valid_parenteses(""))         // false
	log.Println(valid_parenteses("{)"))       // false
	log.Println(valid_parenteses("{}()"))     // true
}

func valid_parenteses(s string) bool {
	if s == "" {
		return false
	}
	
	var stack []rune
	m := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	for _, symbol := range s {
		if _, ok := m[string(symbol)]; ok {

			if len(stack) > 0 && m[string(symbol)] == string(stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}

		} else {
			stack = append(stack, symbol)
		}
	}

	return len(stack) == 0
}
