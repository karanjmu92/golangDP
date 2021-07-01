package main

import (
	"fmt"
)

func main() {
	fmt.Println(gridTraveller(1, 1))
	fmt.Println(gridTraveller(2, 3))
	fmt.Println(gridTraveller(3, 2))
	fmt.Println(gridTraveller(3, 3))
	fmt.Println(gridTraveller(18, 18))
}

func gridTraveller(m, n int) int {

	table := make([][]int, m+1)
	for i, _ := range table {
		table[i] = make([]int, n+1)
	}

	table[1][1] = 1

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			current := table[i][j]
			if j < n {
				table[i][j+1] += current
			}
			if i < m {
				table[i+1][j] += current
			}

		}
	}
	return table[m][n]

}
