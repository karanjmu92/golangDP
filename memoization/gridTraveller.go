package main

import (
	"fmt"
	"strconv"
)

func gridTraveller(m, n int, memo map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if val, ok := memo[key]; ok {
		return val
	}

	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	memo[key] = gridTraveller(m-1, n, memo) + gridTraveller(m, n-1, memo)
	return memo[key]

}

func main() {
	mp := make(map[string]int)
	fmt.Println(gridTraveller(2, 3, mp))
	fmt.Println(gridTraveller(3, 2, mp))
	fmt.Println(gridTraveller(3, 3, mp))
	fmt.Println(gridTraveller(5, 4, mp))
	fmt.Println(gridTraveller(18, 18, mp))
}

