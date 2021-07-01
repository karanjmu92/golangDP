package main

import (
	"fmt"
)

func main() {
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}))
}

func countConstruct(target string, wordBank []string) int {

	table := make([]int, len(target)+1)
	table[0] = 1

	for i := 0; i <= len(target); i++ {

		if table[i] > 0 {
			for _, word := range wordBank {
				if i+len(word) < len(table) {
					if target[i:i+len(word)] == word {
						table[i+len(word)] += table[i]

					}
				}

			}
		}
	}
	return table[len(target)]
}
