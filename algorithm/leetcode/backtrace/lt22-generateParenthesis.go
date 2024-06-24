package main

import "fmt"

func backtrackGenPas(left, right, n int, state string, res *[]string) {
	if left == n && right == n {
		*res = append(*res, state)
		return
	}
	if left < n {
		backtrackGenPas(left+1, right, n, state+"(", res)
	}
	if right < n && left > right {
		backtrackGenPas(left, right+1, n, state+")", res)
	}
}

// [L22-中等] 括号生成
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrackGenPas(0, 0, n, "", &res)
	return res
}

func main() {
	fmt.Println(generateParenthesis(2))
}
