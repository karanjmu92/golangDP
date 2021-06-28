package main

import (
	"fmt"
)

func main() {

	fmt.Println(canSum(7, []int{2, 3}, make(map[int]bool)))
	fmt.Println(canSum(7, []int{5, 3, 4, 7}, make(map[int]bool)))
	fmt.Println(canSum(7, []int{2, 4}, make(map[int]bool)))
	fmt.Println(canSum(8, []int{2, 3, 5}, make(map[int]bool)))
	fmt.Println(canSum(300, []int{7, 14}, make(map[int]bool)))
}

func canSum(targetSum int, numbers []int, mp map[int]bool) bool {

	if val, ok := mp[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, num := range numbers {
		remainder := targetSum - num
		if canSum(remainder, numbers, mp) == true {
			mp[targetSum] = true
			return true
		}
	}
	mp[targetSum] = false
	return false
}

