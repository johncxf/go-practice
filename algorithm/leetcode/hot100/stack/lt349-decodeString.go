// [L349-中等] 字符串解码
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	// 栈，存储数字、左括号、字母
	stack := make([]string, 0)
	// 遍历字符串
	for i := 0; i < len(s); {
		cur := s[i]
		if cur >= '0' && cur <= '9' { // 数字
			// 获取连续数字（可能是2位数或者3位数）
			digits := getDigits(s, &i)
			// 将数字入栈
			stack = append(stack, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' { // 字母或者左括号
			// 入栈
			stack = append(stack, string(cur))
			i++
		} else { // 右括号
			sub := make([]string, 0)
			// 字母出栈，直到遇到左括号停止
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 反转字母序列
			for j := 0; j < len(sub)/2; j++ {
				sub[j], sub[len(sub)-j-1] = sub[len(sub)-j-1], sub[j]
			}
			// 左括号出栈
			stack = stack[:len(stack)-1]
			// 取出栈顶元素（数字），并转化成 int 类型，这个就是前面取出字符串的倍数
			repeat, _ := strconv.Atoi(stack[len(stack)-1])
			// 数字出栈
			stack = stack[:len(stack)-1]
			// 将字符串数组转化成字符串，并生成重复的字符串
			t := strings.Repeat(strings.Join(sub, ""), repeat)
			// 字符串入栈
			stack = append(stack, t)
			i++
		}
	}
	return strings.Join(stack, "")
}

// 获取连续的数字
func getDigits(s string, i *int) string {
	ret := ""
	for ; s[*i] >= '0' && s[*i] <= '9'; *i++ {
		ret += string(s[*i])
	}
	return ret
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}
