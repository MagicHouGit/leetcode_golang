// 54 螺旋矩阵

// 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

// 示例 1:

// 输入:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// 输出: [1,2,3,6,9,8,7,4,5]
// 示例 2:

// 输入:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// 输出: [1,2,3,4,8,12,11,10,9,5,6,7]

package main

import "fmt"

func main() {
	// tb := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// tb := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	// tb := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	// tb := [][]int{{1}}
	// tb := [][]int{{1, 2}, {3, 4}}
	// tb := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	// tb := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	tb := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println(spiralOrder(tb))
}

// 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗 :2 MB, 在所有 Go 提交中击败了73.15%的用户
func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	var res []int
	if n == 0 {
		return res
	}
	l := len(matrix[0])
	//order 从1开始循环，x y的变化量
	r := []int{-1, 1, 1, -1}
	//确定下一次order的起始位置，的变化量
	r2 := []int{1, 1, -1, -1}
	x, y := 0, 0
	var Or int
	//每一圈有4个order，
	var order int
	if l > n {
		order = 2*n - 1
	} else {
		order = 2 * l
	}
	for i := 1; i <= order; i++ {
		Or = i % 4
		if i%2 == 1 {
			for j := 0; j < l-(i/2); j++ {
				res = append(res, matrix[x][y])
				y += r[Or]
			}
			y = y - r[Or]
			x += r2[Or]
		} else {
			for j := 0; j < n-(i/2); j++ {
				res = append(res, matrix[x][y])
				x += r[Or]
			}
			x = x - r[Or]
			y += r2[Or]
		}
	}
	return res
}
