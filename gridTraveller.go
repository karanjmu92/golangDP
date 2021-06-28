package main

import (
	"fmt"
	"strconv"
)

func gridTraveller(m, n int, hm map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if val, ok := hm[key]; ok {
		return val
	}

	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	hm[key] = gridTraveller(m-1, n, hm) + gridTraveller(m, n-1, hm)
	return hm[key]

}

func main() {
	mp := make(map[string]int)
	fmt.Println(gridTraveller(2, 3, mp))
	fmt.Println(gridTraveller(3, 2, mp))
	fmt.Println(gridTraveller(3, 3, mp))
	fmt.Println(gridTraveller(5, 4, mp))
	fmt.Println(gridTraveller(18, 18, mp))
}

