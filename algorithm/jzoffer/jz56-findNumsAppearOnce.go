package main

import "fmt"

/**
 * [JZ56-中等] 数组中只出现一次的两个数字
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func findNumsAppearOnce(array []int) []int {
	hash := map[int]int{}

	// 统计每个数出现的次数
	for _, v := range array {
		if _, ok := hash[v]; ok {
			hash[v] += 1
		} else {
			hash[v] = 1
		}
	}

	// 遍历查找次数等于 1 的数
	result := []int{}
	for k, v := range hash {
		if v == 1 {
			result = append(result, k)
		}
	}

	// 排序 - 就两个数，判断下大小即可
	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			result[i], result[i+1] = result[i+1], result[i]
		}
	}

	return result
}

func main() {
	ret := findNumsAppearOnce([]int{1, 2, 3, 3, 2, 9})
	fmt.Println(ret)
}
