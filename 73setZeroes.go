// 73. 矩阵置零
// 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

// 示例 1:

// 输入:
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// 输出:
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]
// 示例 2:

// 输入:
// [
//   [0,1,2,0],
//   [3,4,5,2],
//   [1,3,1,5]
// ]
// 输出:
// [
//   [0,0,0,0],
//   [0,4,5,0],
//   [0,3,1,0]
// ]
// 进阶:

// 一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个常数空间的解决方案吗？
package main

import "fmt"

func main() {
	// var test [][]int{}
	test := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(test)
	setZeroes(test)
	fmt.Println(test)
}

// 执行用时：8 ms, 在所有 Go 提交中击败了99.42%的用户
// 内存消耗：6.2 MB, 在所有 Go 提交中击败了30.42%的用户
func setZeroes(matrix [][]int) {
	// res := [][]int{}
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	zeroRow := make([]int, row)
	zeroCol := make([]int, col)
	for i := 0; i < row; i++ {
		for ii := 0; ii < col; ii++ {
			if matrix[i][ii] == 0 {
				zeroRow[i] = 1
				zeroCol[ii] = 1
			}
		}
	}
	for i := 0; i < row; i++ {
		for ii := 0; ii < col; ii++ {
			if zeroRow[i] == 1 {
				matrix[i][ii] = 0
			}
			if zeroCol[ii] == 1 {
				matrix[i][ii] = 0
			}

		}

	}
	// return res

}
