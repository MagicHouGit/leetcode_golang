// 118. 杨辉三角
// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

// 在杨辉三角中，每个数是它左上方和右上方的数的和。

// 示例:

// 输入: 5
// 输出:
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

package main

import "fmt"

func main() {
	x := 8
	fmt.Println(generate(x))
}
func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}
	// res[0]= {1}
	res = append(res, []int{1})
	// l := (numRows - 1) / 2

	// if l == 0{
	// 	return res = append{}
	// }
	for i := 1; i < numRows; i++ {
		// temp := []int{}
		temp := make([]int, i+1, i+1)
		temp[0], temp[i] = 1, 1
		l := i / 2
		res = append(res, temp)
		for j := 1; j <= l && i > 1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
			res[i][i-j] = res[i-1][j-1] + res[i-1][j]
			// res = append(res)
		}
	}
	return res
}
