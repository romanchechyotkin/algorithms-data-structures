package main

import "fmt"

func main() {
	fmt.Println(isValid("({[{}[]]})"))
	fmt.Println(isValid("({{]}})"))
	fmt.Println(isValid("({{[}})"))
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	var stack []rune 
	closed := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	for _, ch := range s {
		if _, ok := closed[string(ch)]; ok {
			if len(stack) > 0 && closed[string(ch)] == string(stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}

		} else {
			stack = append(stack, ch)	
		}
	}


	return len(stack) == 0
}