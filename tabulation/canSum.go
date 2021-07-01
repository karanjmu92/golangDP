package main

import (
	"fmt"
)

func main() {
	fmt.Println(canSum(7, []int{2, 3}))
	fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	fmt.Println(canSum(7, []int{2, 4}))
	fmt.Println(canSum(8, []int{2, 3, 5}))
	fmt.Println(canSum(300, []int{7, 14}))
}

func canSum(targetSum int, numbers []int) bool {

	table := make([]bool, targetSum+1)
	table[0] = true
	for i := 0; i <= targetSum; i++ {
		if table[i] {

			for _, num := range numbers {
				if i+num <= targetSum {
					table[i+num] = true
				}
			}

		}
	}
	return table[targetSum]

}
