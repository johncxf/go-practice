package main

import (
	"fmt"
	//"math"
)

var variableType = func() {
	var v1 int            // 整型
	var v2 string         // 字符串
	var v3 bool           // 布尔型
	var v4 [10]int        // 数组，数组元素类型为整型
	var v5 struct {       // 结构体，成员变量 f 的类型为64位浮点型
		f float64
	}
	var v6 *int             // 指针，指向整型
	var v7 map[string]int   // map（字典），key为字符串类型，value为整型
	var v8 func(a int) int  // 函数，参数类型为整型，返回值类型为整型
	fmt.Println("各种类型变量默认值")
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8)
}

/**
 * 【暴力】数组中重复的数字
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func duplicate(numbers []int) int {
	ret := -1
	// 遍历数组
	for i := 0; i < len(numbers); i++ {
		//fmt.Println("Element", i, "of arr is", numbers[i])
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if numbers[i] == numbers[j] {
				ret = numbers[i]
			}
		}
	}
	return ret
}

/**
 * 构建乘积数组
 *
 * @param A int整型一维数组
 * @return int整型一维数组
 */
func multiply( A []int ) []int {
	n := len(A)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		B[i] = 1;
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			B[i] *= A[j]
		}
	}
	return B
}

func multiplyB( A []int ) []int {
	n := len(A)
	letf := 1
	B := make([]int, n)
	for i, v := range A {
		B[i] = letf
		letf *= v
	}
	right := 1
	for j := n - 1; j >= 0; j-- {
		B[j] *= right
		right *= A[j]
	}
	return B
}

/**
 * 不用加减乘除做加法
 *
 * @param num1 int整型
 * @param num2 int整型
 * @return int整型
 */
func add(num1 int, num2 int) int {
	for num2 != 0 {
		c := (num1 & num2) << 1		//进位
		num1 ^= num2				//非进位和
		num2 = c
	}
	return num1
}

/**
 * 替换字符串
 *
 * @param s string字符串
 * @return string字符串
 */
func replaceSpace(s string) string {
	res := ""
	for _, v := range s {
		if v == ' ' {
			res += "%20"
		} else {
			res += string(v)
		}
	}
	return res
}

/**
 * 第一个只出现一次的字符串
 *
 * @param str string字符串
 * @return int整型
 */
func firstNotRepeatingChar(str string) int {
	if len(str) == 0 {
		return -1
	}
	keyMap := [256]int {0}
	for i := range str {
		keyMap[str[i]]++
	}
	for i , v := range str {
		if keyMap[v] == 1 {
			return i
		}
	}
	return -1
}

/**
 * 旋转数组的最小数字
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	// 边界值处理
	if len(rotateArray) == 0 {
		return 0
	}
	if len(rotateArray) == 1 {
		return rotateArray[0]
	}

	// 初始化中间变量
	var mid int
	// 定义左右变量
	left, right := 0, len(rotateArray) - 1
	for rotateArray[left] >= rotateArray[right] {
		// 退出条件
		if (right - left) == 1 {
			mid = right
			break
		}
		// 中间变量
		mid = (left + right) / 2
		if rotateArray[mid] > rotateArray[right] {// 中间值大于最右侧值，说明最小值在右侧
			left = mid
		} else if rotateArray[mid] < rotateArray[right] {// 中间值小于最右侧值，说明最小值在左侧
			right = mid
		} else {
			right--
		}
	}
	return rotateArray[mid]
}

// 定义节点
type ListNode struct {
	Val int
	Next *ListNode
}

//var headNode = new(ListNode)
//
//// 添加节点
//func addListNode(t *Node, v int) int {
//	if head == nil {
//		t = &Node{v, nil}
//		head = t
//		return 0
//	}
//	if v == t.Value {
//		fmt.Println("节点已存在:", v)
//		return -1
//	}
//	// 如果当前节点下一个节点为空
//	if t.Next == nil {
//		t.Next = &Node{v, nil}
//		return -2
//	}
//	// 如果当前节点下一个节点不为空
//	return addNode(t.Next, v)
//}

func main() {
	//arr := []int {2, 6, 1, 0, 2, 5, 3}
	//res := duplicate(arr)

	//arr1 := []int {1,2,3,4,5}
	//res := multiply(arr1)

	//num1 := 1
	//num2 := 2
	//res := add(num1, num2)

	//addNode(head, 6)
	//
	//res := printListFromTailToHead(head)
	//fmt.Println(res)
}