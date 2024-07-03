// [L739-中等] 每日温度
package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		// 当前温度
		temperature := temperatures[i]
		// 栈不为空，并且栈顶元素下标对应温度大于当前温度
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			// 栈顶元素出栈
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 更新答案
			ans[index] = i - index
		}
		// 当前温度下标入栈
		stack = append(stack, i)
	}
	return ans
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
