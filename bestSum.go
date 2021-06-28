package main

import (
	"fmt"
)

func main() {

	fmt.Println(bestSum(7, []int{5, 3, 4, 7}, make(map[int][]int)))
	fmt.Println(bestSum(8, []int{2, 3, 5}, make(map[int][]int)))
	fmt.Println(bestSum(8, []int{1, 4, 5}, make(map[int][]int)))
	fmt.Println(bestSum(100, []int{1, 2, 5, 25}, make(map[int][]int)))
}

func bestSum(targetSum int, numbers []int, mp map[int][]int) []int {

	if val, ok := mp[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	var shortestCombination []int = nil
	for _, num := range numbers {
		remainder := targetSum - num
		remainderCombination := bestSum(remainder, numbers, mp)
		if remainderCombination != nil {
			combination := append(remainderCombination, num)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}

	mp[targetSum] = shortestCombination
	return shortestCombination

}

