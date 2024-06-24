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

func backtrackLetter(digits string, index int, combination string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, combination)
		return
	}
	digit := string(digits[index])
	letters := phoneMap[digit]
	for i := 0; i < len(letters); i++ {
		backtrackLetter(digits, index+1, combination+string(letters[i]), res)
	}
}

// [L17-中等] 电话号码的字母组合
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	backtrackLetter(digits, 0, "", &res)
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
