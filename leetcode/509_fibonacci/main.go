package main

import "log"

func main() {
	log.Println(rec_fib(10) == fib(10))
	log.Println(rec_fib(2) == fib(2))
	log.Println(rec_fib(3) == fib(3))
	log.Println(rec_fib(4) == fib(4))

	log.Println(factorial(4) == rec_factorial(4))
	log.Println(factorial(5) == rec_factorial(5))
}

func rec_fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return rec_fib(n-2) + rec_fib(n-1)
}

func fib(n int) int {
	a := 0
	b := 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func factorial(n int) int {
	sum := 1

	for i := 1; i <= n; i++ {
		sum = sum * i
	}

	return sum
}
func rec_factorial(n int) int {
	if n == 1 {
		return n
	}

	return n * factorial(n-1)
}
