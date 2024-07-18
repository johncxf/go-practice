// [L210-中等] 课程表II
package main

import "fmt"

// 深度优先搜索（dfs）
func findOrderDfs(numCourses int, prerequisites [][]int) []int {
	// 结果集
	res := make([]int, 0)
	// 构建邻接表（存储有向图）
	// i 对应于一个节点（即课程），edges[i] 则是一个包含所有以课程 i 为先修课程的课程列表。
	edges := make([][]int, numCourses)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	// 栈结构：标记每个节点的状态：0=未搜索，1=搜索中，2=已完成
	visited := make([]int, numCourses)
	// 是否存在拓扑排序（如果一个有向图包含环，则不存在）
	valid := true
	// 深度优先遍历
	var dfs func(u int)
	dfs = func(u int) {
		// 标记为搜索中：1
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 { // 未搜索，则开始搜索
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 { // 搜索中，说明找到了环
				valid = false
				return
			}
		}
		// 搜索完成，修改为已完成状态
		visited[u] = 2
		// 记录结果
		res = append(res, u)
	}
	// 遍历，对未搜索的节点进行深度优先搜索
	// 为什么遍历 numCourses？因为题目中说到 0 ～ numCourses-1 为对应要选的课程
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	// 非拓扑排序，则不存在，返回空数组
	if !valid {
		return []int{}
	}
	// 反转结果
	for i := 0; i < len(res)/2; i++ {
		res[i], res[numCourses-i-1] = res[numCourses-i-1], res[i]
	}
	return res
}

// 广度优先搜索（bfs）
func findOrderBfs(numCourses int, prerequisites [][]int) []int {
	// 结果集
	res := make([]int, 0)
	// 每门课程的入度（即有多少课程依赖它）
	inDegree := make([]int, numCourses)
	// 构建邻接表
	edges := make([][]int, numCourses)
	for _, info := range prerequisites {
		// edges[i] 表示课程 i 的所有后继课程（即依赖于课程 i 的课程）。
		edges[info[1]] = append(edges[info[1]], info[0])
		// 获取每门课程的入度（有多少课程依赖它）
		inDegree[info[0]]++
	}
	// 先找出所有入度为0的节点（即先找出不依赖其他课程的课程）
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 遍历入度为0的节点
	for len(queue) > 0 {
		// 出队
		u := queue[0]
		queue = queue[1:]
		// 入度为0，则说明已不依赖其他节点，加入结果
		res = append(res, u)
		// 删除依赖当前课程的度（边）
		for _, v := range edges[u] {
			inDegree[v]--
			// 度为0的时候则可以加入 queue 队列中
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	// 解雇长度不等于课程数，说明有无法删除的节点，则说明存在环
	if len(res) != numCourses {
		return []int{}
	}
	return res
}

func main() {
	fmt.Println("深度优先搜索（dfs）:")
	fmt.Println(findOrderDfs(2, [][]int{{1, 0}}))
	fmt.Println(findOrderDfs(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))

	fmt.Println("广度优先搜索（bfs）:")
	fmt.Println(findOrderBfs(2, [][]int{{1, 0}}))
	fmt.Println(findOrderBfs(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
