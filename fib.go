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

func fib(n int, hm map[int]int) int{

	if val, ok := hm[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}

	hm[n] = fib(n-1, hm) + fib(n-2, hm)
	return hm[n]

}
