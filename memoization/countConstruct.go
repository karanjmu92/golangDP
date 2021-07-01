package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, make(map[string]int)))
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, make(map[string]int)))
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, make(map[string]int)))
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, make(map[string]int)))
	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, make(map[string]int)))
}

func countConstruct(target string, wordBank []string, memo map[string]int) int {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return 1
	}
	totalCount := 0
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {

			remainder := strings.TrimPrefix(target, word)
			numWaysForRest := countConstruct(remainder, wordBank, memo)
			totalCount += numWaysForRest
		}
	}
	memo[target] = totalCount
	return totalCount

}