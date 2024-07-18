// [L56-中等] 合并区间
package main

import (
	"fmt"
	"sort"
)

// 排序+比较
func mergeInterval1(intervals [][]int) [][]int {
	// 先根据数组左侧大小对二维数组进行从小到大排序
	n := len(intervals)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	// 遍历数组
	res := make([][]int, 0)
	pre := intervals[0]
	for i := 1; i < n; i++ {
		// 重合：2个条件
		// 1、前一个数组的右区间大于下一个数组的左区间
		if pre[1] >= intervals[i][0] {
			// 2、前一个数组的右区间大于下一个数组的右区间
			if pre[1] < intervals[i][1] {
				pre[1] = intervals[i][1]
			}

		} else {
			res = append(res, pre)
			pre = intervals[i]
		}
	}
	res = append(res, pre)
	return res
}

// 排序+比较
func mergeInterval2(intervals [][]int) [][]int {
	// 先根据数组左侧大小对二维数组进行从小到大排序
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 遍历数组
	res := make([][]int, 0)
	pre := intervals[0]
	for i := 1; i < n; i++ {
		// 重合：2个条件
		// 1、前一个数组的右区间大于下一个数组的左区间
		if pre[1] >= intervals[i][0] {
			// 2、前一个数组的右区间大于下一个数组的右区间
			if pre[1] < intervals[i][1] {
				pre[1] = intervals[i][1]
			}

		} else {
			res = append(res, pre)
			pre = intervals[i]
		}
	}
	res = append(res, pre)
	return res
}

func main() {
	fmt.Println(mergeInterval1([][]int{{1, 3}, {15, 18}, {2, 6}, {8, 10}}))

	fmt.Println(mergeInterval2([][]int{{1, 3}, {15, 18}, {2, 6}, {8, 10}}))

}
