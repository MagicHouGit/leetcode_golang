//494 对角线遍历
// 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。

// 示例:

// 输入:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]

// 输出:  [1,2,4,7,5,3,6,8,9]
// 解释:
// 说明:
// 给定矩阵中的元素总数不会超过 100000 。

package main

// func main() {
// 	// tmatrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
// 	tmatrix := [][]int{{1}}
// 	// tmatrix := [][]int{{1, 4}, {5, 8}, {9, 12}, {15, 16}}
// 	// tmatrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
// 	fmt.Println(findDiagonalOrder(tmatrix))
// }

func findDiagonalOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	//false 非x ，true 代表x
	// var xy bool
	// xy = true
	pointX := 0
	pointY := 0
	//升降
	gdX := 0
	gdY := 0
	x := len(matrix)
	y := len(matrix[0])
	//斜边条数
	fd := x + y - 1
	// if fd < 0 {
	// 	return result
	// }
	var count []int
	min, max := 0, 0
	if x > y {
		max = x
		min = y
	} else {
		min = x
		max = y
	}
	for i := 1; i < min; i++ {
		// count[i] = i + 1
		count = append(count, i)
	}
	for i := 0; i < max-min+1; i++ {
		count = append(count, min)
	}
	for i := min; i > 1; i-- {
		count = append(count, i-1)
	}

	// result = append(result, matrix[pointX][pointY])
	for i := 0; i < fd; i++ {

		result = append(result, matrix[pointX][pointY])
		// if count[i] != min {
		// 	// xy = ^xy
		// 	// xy != xy
		// 	if xy == true {
		// 		xy = false
		// 	} else {
		// 		xy = true
		// 	}
		// }
		for j := 0; j < count[i]-1; j++ {
			pointX += gdX
			pointY += gdY
			result = append(result, matrix[pointX][pointY])
		}
		if i%2 == 0 {
			gdX = 1
			gdY = -1
			if pointY == y-1 {
				pointX++
			} else {
				pointY++
			}

		} else {
			gdX = -1
			gdY = 1
			if pointX == x-1 {
				pointY++
			} else {
				pointX++
			}

		}
		// if xy == true {
		// 	pointX++
		// } else {
		// 	pointY++
		// }
	}
	// result = count
	return result
}
