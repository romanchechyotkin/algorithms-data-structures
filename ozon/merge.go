package main

import (
	"log"
	"sync"
)

func merge(arrs ...[]int) []int {
	l := 0
	for _, arr := range arrs {
		l += len(arr)
	}

	res := make([]int, l)
	h := res[0:l]
	log.Println(h, len(h), cap(h))
	for _, arr := range arrs {
		copy(h, arr)
		log.Println(h, len(h), cap(h))

		h = h[len(arr):]
		log.Println(h, len(h), cap(h))
	}

	return res
}

func concurrent_merge(arrs ...[]int) []int {
	if len(arrs) == 0 {
		return nil
	}

	res := make([]int, 0, len(arrs))
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < len(arrs); i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for _, val := range arrs[i] {
				mu.Lock()
				res = append(res, val)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	return res
}

func main() {
	log.Println(merge([]int{1, 2}, []int{3, 4}, []int{5, 6}))
	log.Println(merge([]int{1, 2}, []int{5, 6}))
	log.Println(merge())

	log.Println(concurrent_merge([]int{1, 2}, []int{3, 4}, []int{5, 6}))
	log.Println(concurrent_merge([]int{1, 2}, []int{5, 6}))
	log.Println(concurrent_merge())
}
