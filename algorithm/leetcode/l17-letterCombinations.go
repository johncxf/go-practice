package main

import (
	"fmt"
)

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combinations = []string{}
	dfs(digits, 0, "")
	return combinations
}

func dfs(digits string, index int, combination string) {
	if len(digits) == index {
		combinations = append(combinations, combination)
		return
	}

	digit := string(digits[index])
	letters := phoneMap[digit]
	for i := 0; i < len(letters); i++ {
		dfs(digits, index+1, combination+string(letters[i]))
	}
}

func main() {
	fmt.Println(letterCombinations("2"))
}
