package main

import (
	"fmt"
)

func main() {

	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))
	fmt.Println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(allConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(allConstruct("eeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}))
}

func allConstruct(target string, wordBank []string) [][]string {

	table := make([][][]string, len(target)+1)
	table[0] = [][]string{{}}

	for i := 0; i <= len(target); i++ {
		if len(table[i]) > 0 {
			for _, word := range wordBank {
				if i+len(word) < len(table) {
					if target[i:i+len(word)] == word {
						newCombination := make([][]string, len(table[i]))
						for j, val := range table[i] {
							newCombination[j] = append(val, word)
						}

						table[i+len(word)] = append(table[i+len(word)], newCombination...)

					}
				}

			}
		}
	}

	return table[len(target)]
}
