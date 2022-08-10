package main

import (
	"fmt"
)

func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}

	bracketsMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []byte{}
	for i := 0; i < length; i++ {
		if bracketsMap[s[i]] != 0 {
			stack = append(stack, bracketsMap[s[i]])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
