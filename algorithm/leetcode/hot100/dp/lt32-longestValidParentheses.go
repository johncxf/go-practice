// [L32-困难] 最长有效括号
package main

import "fmt"

// 用一个标记数组存放每个位置是否匹配信息
func longestValidParentheses1(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	// 栈，存储左括号下标，用于判断是否匹配成功
	stack := make([]int, 0)
	// 标记数组，标记字符串该位置是否匹配成功
	flags := make([]bool, n)
	// 遍历字符串，进行标记
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				continue
			} else {
				// 左括号下标
				leftIndex := stack[len(stack)-1]
				// 出栈
				stack = stack[:len(stack)-1]
				// 标记
				flags[i], flags[leftIndex] = true, true
			}
		}
	}
	// 遍历标记数组，找出连续匹配成功的最大字符串
	max := 0
	cur := 0
	for _, flag := range flags {
		if flag {
			cur++
		} else {
			if cur > max {
				max = cur
			}
			cur = 0
		}
	}
	// 最后一个字符也匹配成功的情况
	if cur > max {
		max = cur
	}
	return max
}

// 在上面基础上进一步优化
func longestValidParentheses2(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	// 栈，栈低存储没有被匹配的最后一个右括号，其他元素为左括号
	stack := make([]int, 0)
	// 初始化一个 -1
	stack = append(stack, -1)
	max := 0
	// 遍历字符串，进行标记
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 出栈
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 栈为空，下标入栈，此时刚好是栈低元素
				stack = append(stack, i)
			} else {
				// 不为空，更新最大值
				if i-stack[len(stack)-1] > max {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses1("(()"))
	fmt.Println(longestValidParentheses1(")()())"))
	fmt.Println(longestValidParentheses1("()(()"))

	fmt.Println(longestValidParentheses2("(()"))
	fmt.Println(longestValidParentheses2(")()())"))
	fmt.Println(longestValidParentheses2("()(()"))
}
