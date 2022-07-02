package main

import "fmt"

/**
 * [JZ61-简单] 扑克牌顺子
 *
 * @param numbers int整型一维数组
 * @return bool布尔型
 */
func isContinuous(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == 0 || numbers[j] == 0 {
				continue
			}
			tmp := numbers[i] - numbers[j]
			// 存在两数相等或者相差大于 4，则肯定不连续
			if tmp > 4 || tmp < -4 || tmp == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	ret := isContinuous([]int{6, 0, 2, 0, 4})
	fmt.Println(ret)
}
