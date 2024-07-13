package main

import "fmt"

// [L56-中等] 合并区间
func mergeInterval(intervals [][]int) [][]int {
	// 先安装数组左端元素从小到大进行排序
	n := len(intervals)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	res := make([][]int, 0)
	tmp := intervals[0]
	for i := 1; i < n; i++ {
		if tmp[1] < intervals[i][0] { // 没有重合
			res = append(res, tmp)
			tmp = intervals[i]
		} else { // 重合
			if tmp[1] < intervals[i][1] {
				tmp[1] = intervals[i][1]
			}
		}
	}
	res = append(res, tmp)
	return res
}

func main() {
	fmt.Println(mergeInterval([][]int{{1, 3}, {15, 18}, {2, 6}, {8, 10}}))
}
