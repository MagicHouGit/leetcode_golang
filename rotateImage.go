//旋转图像
package main

import "fmt"

func main() {
	tmatrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	for _, v := range tmatrix {
		for _, va := range v {
			fmt.Printf("[%d]", va)
		}
		fmt.Printf("\n")
	}
	rotate(tmatrix)

	for _, v := range tmatrix {
		for _, va := range v {
			fmt.Printf("[%d]", va)
		}
		fmt.Printf("\n")
	}
}

// func rotate(matrix [][]int) {
// 	// tempOne := matrix[0][0]
// 	//执行用时 : 4 ms, 在Rotate Image的Go提交中击败了94.12% 的用户
// 	//内存消耗 : 2.7 MB, 在Rotate Image的Go提交中击败了45.90% 的用户
// 	//这题我敢保障不出一个月我就会忘了我的结题思路，所以 结题思路见ipad上的paper
// 	tempOne := 0
// 	size := len(matrix[0])
// 	// time := 0
// 	layer := 0
// 	layerNum := 0
// 	// fmt.Printf("%d\n", size)
// 	// time = ((size+1)/2)*((size+1)/2) - ((size + 1) / 2) + 1
// 	layer = size / 2
// 	for i := 1; i <= layer; i++ {
// 		//每层的个数f(size-1)-(layer-1)*2
// 		layerNum = (size - 1) - (i-1)*2
// 		for j := 0; j < layerNum; j++ {
// 			tempOne = matrix[i-1][i-1+j]
// 			matrix[i-1][i-1+j] = matrix[size-i-j][i-1]
// 			matrix[size-i-j][i-1] = matrix[size-i][size-i-j]
// 			matrix[size-i][size-i-j] = matrix[i-1+j][size-i]
// 			matrix[i-1+j][size-i] = tempOne
// 		}
// 	}
// }
//老规矩从leecode上抄来的看看
func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 0 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// for i := 0; i < n; i++ {
	// 	for j, k := 0, n-1; j < k; {
	// 		matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
	// 		j++
	// 		k--
	// 	}
	// }
}
