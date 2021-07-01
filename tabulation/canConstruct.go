package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}))
}

func canConstruct(target string, wordBank []string) bool {

	table := make([]bool, len(target)+1)
	table[0] = true

	for i := 0; i <= len(target); i++ {

		if table[i] {
			for _, word := range wordBank {
				if i+len(word) < len(table) {
					if target[i:i+len(word)] == word {
						table[i+len(word)] = true

					}
				}

			}
		}
	}
	return table[len(target)]
}
