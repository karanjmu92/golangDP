package main

import (
	"fmt"
)

func main() {

	fmt.Println(howSum(7, []int{2, 3}, make(map[int][]int)))
	fmt.Println(howSum(7, []int{5, 3, 4, 7}, make(map[int][]int)))
	fmt.Println(howSum(7, []int{2, 4}, make(map[int][]int)))
	fmt.Println(howSum(8, []int{2, 3, 5}, make(map[int][]int)))
	fmt.Println(howSum(300, []int{7, 14}, make(map[int][]int)))
}

func howSum(targetSum int, numbers []int, mp map[int][]int) []int {

	if val, ok := mp[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := targetSum - num
		remainderResult := howSum(remainder, numbers, mp)
		if remainderResult != nil {
			mp[targetSum] = append(remainderResult, num)
			return mp[targetSum]
		}
	}
	mp[targetSum] = nil
	return nil

}

