// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：
// 输入：nums = [1]
// 输出：[[1]]
package main

import (
	"fmt"
	"net/http"

	// _ "net/http/pprof"
	_ "runtime/pprof"
)

func main() {
	// testNum := []int{1, 2, 3}
	// testNum := []int{1, 2}
	// testNum := []int{0, 1}
	// testNum := []int{1}
	testNum := []int{1, 2, 3, 4}
	// testNum1 := []int{1}
	fmt.Println(testNum)
	fmt.Println(permute(testNum))
	// fmt.Println(chaban(7, 1, testNum1))
	// go func() {
	// 	err := http.ListenAndServe(":6060", nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()
	// go func() {
	// for {

	// log.Println(http.ListenAndServe("localhost:6060", nil))
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	// }
	// }()
}

// 插板法，

// 这晦涩难懂的东西我都不知道该怎么写注释，
// 2021-7-18，这代码明天我看看不懂了
func permute(nums []int) [][]int {
	var res [][]int
	res = append(res, append([]int(nil), nums[0]))
	if len(nums) == 1 {
		return res
	}
	// l 1,2,3,4
	//   0,1,2,3
	cont := 1
	l := len(nums)
	for i := 1; i < l; i++ {
		cont = cont * i
		for ii := 0; ii < cont; ii++ {
			for iii := 0; iii < i+1; iii++ {
				res = append(res, chaban(nums[i], iii, res[0]))
			}
			res = res[1:]
		}
	}
	return res
}

// c 要插的数字，site 要插的位置，被插的数组
func chaban(c int, site int, sle []int) []int {
	l := len(sle)
	if site == l+1 {
		sle = append(sle, c)
	}
	sle = append(sle, sle[l-1])
	for i := l; i > site; i-- {
		sle[i] = sle[i-1]
	}
	if site < len(sle) {
		sle[site] = c
	}
	//出现地址重复，24组【4，4，4，4】
	res := make([]int, len(sle))
	copy(res, sle)
	return res
}

// func permute(nums []int) [][]int {
// 	if len(nums) == 1 {
// 		return [][]int{nums}
// 	}

// 	res := [][]int{}

// 	for i, num := range nums {
// 		// 把num从 nums 拿出去 得到tmp
// 		tmp := make([]int, len(nums)-1)
// 		copy(tmp[0:], nums[0:i])
// 		copy(tmp[i:], nums[i+1:])

// 		// sub 是把num 拿出去后，数组中剩余数据的全排列
// 		sub := permute(tmp)
// 		for _, s := range sub {
// 			res = append(res, append(s, num))
// 		}
// 	}
// 	return res
// }
