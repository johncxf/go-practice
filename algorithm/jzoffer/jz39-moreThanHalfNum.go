package main

import "fmt"

/**
 * [JZ39-简单] 数组中出现次数超过一半的数字（哈希法）
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func moreThanHalfNum(numbers []int) int {
	length := len(numbers)
	if 1 == length {
		return numbers[0]
	}

	hash := map[int]int{}
	for _, v := range numbers {
		if _, ok := hash[v]; ok {
			hash[v] += 1
		} else {
			hash[v] = 1
		}
	}

	for k, v := range hash {
		if v > length/2 {
			return k
		}
	}

	return 0
}

/**
 * [JZ39-简单] 数组中出现次数超过一半的数字（候选法）
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func moreThanHalfNum2(numbers []int) int {
	length := len(numbers)
	if 1 == length {
		return numbers[0]
	}

	cond, count := -1, 0
	for _, v := range numbers {
		if 0 == count {
			cond = v
			count++
		} else {
			if cond == v {
				count++
			} else {
				count--
			}
		}
	}

	count = 0
	for _, v := range numbers {
		if v == cond {
			count++
		}
	}

	if count > length/2 {
		return cond
	} else {
		return 0
	}
}

func main() {
	arr := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	ret := moreThanHalfNum2(arr)
	fmt.Println(ret)
}
