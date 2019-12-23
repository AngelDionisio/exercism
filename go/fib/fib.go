package main

import (
	"fmt"
	"time"
)

func main() {
	newFib := memoize(fib)

	start := time.Now()
	for i := 1; i < 40; i++ {
		fmt.Println(newFib(i))
	}
	timeNotCached := time.Since(start)

	start = time.Now()
	for i := 1; i < 40; i++ {
		fmt.Println(newFib(i))
	}
	timeCached := time.Since(start)

	fmt.Println()
	fmt.Println("timeNotCached", timeNotCached)
	fmt.Println("timeCached", timeCached)
}

func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if cache[n] == 0 {
			cache[n] = f(n)
		}
		return cache[n]
	}
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
