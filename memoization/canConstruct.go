package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, make(map[string]bool)))
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, make(map[string]bool)))
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, make(map[string]bool)))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, make(map[string]bool)))
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {

			remainder := strings.TrimPrefix(target, word)
			if canConstruct(remainder, wordBank, memo) == true {
				memo[target] = true
				return true
			}

		}
	}
	memo[target] = false
	return false

}
