package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, make(map[string][][]string)))
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}, make(map[string][][]string)))
	fmt.Println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, make(map[string][][]string)))
	fmt.Println(allConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, make(map[string][][]string)))
	fmt.Println(allConstruct("eeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, make(map[string][][]string)))
}

func allConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return [][]string{{}}
	}
	var result [][]string
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := strings.TrimPrefix(target, word)
			suffixWays := allConstruct(suffix, wordBank, memo)
			targetways := make([][]string, len(suffixWays))
			for i, val := range suffixWays {
				targetways[i] = append([]string{word}, val...)
			}
			result = append(result, targetways...)
		}

	}
	memo[target] = result
	return result

}
