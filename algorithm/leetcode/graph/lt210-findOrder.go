// [L210-中等] 课程表II
package main

import "fmt"

// 深度优先搜索（dfs）
func findOrderDfs(numCourses int, prerequisites [][]int) []int {
	// 结果集
	res := make([]int, 0)
	// 遍历过的节点，存入栈
	visited := make([]int, numCourses)
	// 邻接表
	edges := make([][]int, numCourses)
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
	// 构建邻接表
	// i 对应于一个节点（即课程），edges[i] 则是一个包含所有以课程 i 为先修课程的课程列表。
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	// 遍历，对未搜索的节点进行深度优先搜索
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
	// 邻接表
	edges := make([][]int, numCourses)
	// 每门课程的入度
	indeg := make([]int, numCourses)
	// 结果集
	res := make([]int, 0)
	// 构建邻接表
	for _, info := range prerequisites {
		// edges[i] 表示课程 i 的所有后继课程（即依赖于课程 i 的课程）。
		edges[info[1]] = append(edges[info[1]], info[0])
		// 获取每门课程的入度（有多少课程依赖它）
		indeg[info[0]]++
	}
	// 找出所有入度为0的节点
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	// 遍历入度为0的节点
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		// 入度为0，则说明已不依赖其他节点，加入结果
		res = append(res, u)
		// 删除节点（删除节点所有度）
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
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
