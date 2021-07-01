package main

import (
	"fmt"
)

func main() {
	fmt.Println(bestSum(7, []int{2, 3}))
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(7, []int{2, 4}))
	fmt.Println(bestSum(8, []int{2, 3, 5}))
	fmt.Println(bestSum(100, []int{7, 14, 2, 25}))
}

func bestSum(targetSum int, numbers []int) []int {

	table := make([][]int, targetSum+1)
	table[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {

			for _, num := range numbers {
				if i+num <= targetSum {
					if table[i+num] == nil || len(table[i])+1 < len(table[i+num]) {
						table[i+num] = append(table[i], num)
					}
				}
			}

		}
	}
	return table[targetSum]

}
