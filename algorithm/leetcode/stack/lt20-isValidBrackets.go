package main

import (
	"fmt"
)

// [L20-简单] 有效的括号
func isValid(s string) bool {
	bracketsMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if bracketsMap[s[i]] > 0 {
			// 左括号，将对应右括号入栈
			stack = append(stack, bracketsMap[s[i]])
		} else { // 右括号，判断栈中是否有匹配
			// 无匹配，返回 false
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			// 有匹配，进行出栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
