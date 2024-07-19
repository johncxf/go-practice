// [L22-中等] 括号生成
package main

import "fmt"

// 回溯
func backtrackGenPas(left, right, n int, state string, res *[]string) {
	// 左右括号数达到要求，加入结果集
	if left == n && right == n {
		*res = append(*res, state)
		return
	}

	// 先添加左括号，左括号添加条件需要 < n
	if left < n {
		backtrackGenPas(left+1, right, n, state+"(", res)
	}
	// 右括号添加条件需要 < n，并且当前左括号数要大于右括号
	if right < n && left > right {
		backtrackGenPas(left, right+1, n, state+")", res)
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrackGenPas(0, 0, n, "", &res)
	return res
}

func main() {
	fmt.Println(generateParenthesis(2))
}
