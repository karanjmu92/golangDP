package main

import (
	"fmt"
)

func main() {
	
	mp := make(map[int]int)
	fmt.Println(fib(6, mp))
	fmt.Println(fib(7, mp))
	fmt.Println(fib(8, mp))
	fmt.Println(fib(50, mp))
}

func fib(n int, memeo map[int]int) int{

	if val, ok := memeo[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}

	memeo[n] = fib(n-1, memeo) + fib(n-2, memeo)
	return memeo[n]

}
