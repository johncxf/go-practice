package main

import (
	"fmt"
	"math"
)

/**
 * 枚举法
 *
 * @param sum int整型
 * @return int整型二维数组
 */
func findContinuousSequence(sum int) [][]int {
	if sum == 0 {
		return [][]int{{}}
	}

	half := math.Ceil(float64(sum) / 2)
	count := int(half)

	result := make([][]int, 0)
	for i := 1; i < count; i++ {
		tmpArr := []int{i}
		tmpSum := i
		for j := i + 1; j <= count; j++ {
			tmpArr = append(tmpArr, j)
			tmpSum += j
			if tmpSum > sum {
				break
			} else if tmpSum == sum {
				result = append(result, tmpArr)
			}
		}
	}
	return result
}

/**
 * 滑动窗口
 *
 * @param sum int 整型
 * @return int 整型二维数组
 */
func findContinuousSequence2(sum int) [][]int {
	if sum == 0 {
		return [][]int{{}}
	}

	result := make([][]int, 0)

	// 定义左右区间初始值为 1、2
	l, r := 1, 2
	for l < r {
		// 区间所有数之和
		tmpSum := (l + r) * (r - l + 1) / 2

		if tmpSum == sum {
			// 等于预期值，将区间处理成一维数组
			tmpArr := []int{}
			for i := l; i <= r; i++ {
				tmpArr = append(tmpArr, i)
			}
			result = append(result, tmpArr)

			// 移动区间左临界值，继续查找下一个
			l++
		} else if tmpSum < sum {
			// 小于预期值，左区间 l 右移一位
			r++
		} else if tmpSum > sum {
			// 大于预期值，右区间 r 左边移一位
			l++
		}
	}
	return result
}

func main() {
	ret := findContinuousSequence2(9)
	fmt.Println(ret)
}
