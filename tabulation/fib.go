package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(50))
}

func fib(n int) int {

	table := make([]int, n+1)
	table[1] = 1
	for i := 0; i <= n-1; i++ {
		table[i+1] += table[i]
		if i < n-1 {
			table[i+2] += table[i]
		}
	}

	return table[n]

}
