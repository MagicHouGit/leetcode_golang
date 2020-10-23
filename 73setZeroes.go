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
