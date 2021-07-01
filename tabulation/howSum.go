package main

import (
	"fmt"
)

func main() {
	fmt.Println(howSum(7, []int{2, 3}))
	fmt.Println(howSum(7, []int{5, 3, 4, 7}))
	fmt.Println(howSum(7, []int{2, 4}))
	fmt.Println(howSum(8, []int{2, 3, 5}))
	fmt.Println(howSum(300, []int{7, 14}))
}

func howSum(targetSum int, numbers []int) []int {

	table := make([][]int, targetSum+1)
	table[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {

			for _, num := range numbers {
				if i+num <= targetSum {
					table[i+num] = append(table[i], num)
				}
			}

		}
	}
	return table[targetSum]

}
