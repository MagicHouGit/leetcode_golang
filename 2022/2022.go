package main

import "fmt"

func main() {
	o1 := []int{1, 2, 3, 4}
	o2 := []int{1, 2, 3, 4}
	o3 := []int{1, 2, 3, 4}
	o4 := []int{1, 2, 3, 4}
	m1 := 2
	m2 := 2
	m3 := 1
	m4 := 4
	n1 := 2
	n2 := 3
	n3 := 1
	n4 := 1
	fmt.Println(construct2DArray(o1, m1, n1))
	fmt.Println(construct2DArray(o2, m2, n2))
	fmt.Println(construct2DArray(o3, m3, n3))
	fmt.Println(construct2DArray(o4, m4, n4))

}
func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = original[i*n : i*n+n]
	}
	return res
}

// 执行用时：112 ms, 在所有 Go 提交中击败了55.59%的用户
// 内存消耗：9.8 MB, 在所有 Go 提交中击败了24.12%的用户
// func construct2DArray(original []int, m int, n int) [][]int {
// 	res := [][]int{}
// 	if len(original) > m*n || len(original) < m*n {
// 		return res
// 	}
// 	for i := 0; i < m; i++ {
// 		temp := []int{}
// 		for ii := 0; ii < n; ii++ {
// 			temp = append(temp, original[i*n+ii])
// 		}
// 		res = append(res, temp)
// 	}
// 	return res
// }
