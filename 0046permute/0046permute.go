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

	// _ "net/http/pprof"
	_ "runtime/pprof"
)

func main() {
	testNum := []int{1, 2, 3}
	// testNum := []int{1, 2}
	// testNum := []int{0, 1}
	// testNum := []int{1}
	// testNum := []int{1, 2, 3, 4}
	// fmt.Println(testNum)
	fmt.Println(permute(testNum))
	// fmt.Println(chaban(7, 1, testNum1))
	// fmt.Println(permuteOther(testNum))

}

//============================
//2022-02-11
//2022-02-12
func permute(nums []int) [][]int {
	res := [][]int{}
	var backTrack func(subset []int, l int, nums []int)
	backTrack = func(subset []int, l int, nums []int) {
		if len(nums) == 0 {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)

		} else {

			for i := 0; i < l; i++ {
				take := nums[i]
				subset = append(subset, take)
				nums = append(nums[:i], nums[i+1:]...)
				backTrack(subset, len(nums), nums)
				nums = append(nums[:i], append([]int{take}, nums[i:]...)...)
				subset = subset[:len(subset)-1]
			}
		}

	}
	l := len(nums)
	backTrack([]int{}, l, nums)
	return res
}

//=================================
// // 插板法，

// // 这晦涩难懂的东西我都不知道该怎么写注释，
// // 2021-7-18，这代码明天我看看不懂了
// func permute(nums []int) [][]int {
// 	var res [][]int
// 	res = append(res, append([]int(nil), nums[0]))
// 	if len(nums) == 1 {
// 		return res
// 	}
// 	// l 1,2,3,4
// 	//   0,1,2,3
// 	cont := 1
// 	l := len(nums)
// 	for i := 1; i < l; i++ {
// 		cont = cont * i
// 		for ii := 0; ii < cont; ii++ {
// 			for iii := 0; iii < i+1; iii++ {
// 				res = append(res, chaban(nums[i], iii, res[0]))
// 			}
// 			res = res[1:]
// 		}
// 	}
// 	return res
// }

// // c 要插的数字，site 要插的位置，被插的数组
// func chaban(c int, site int, sle []int) []int {
// 	l := len(sle)
// 	if site == l+1 {
// 		sle = append(sle, c)
// 	}
// 	sle = append(sle, sle[l-1])
// 	for i := l; i > site; i-- {
// 		sle[i] = sle[i-1]
// 	}
// 	if site < len(sle) {
// 		sle[site] = c
// 	}
// 	//出现地址重复，24组【4，4，4，4】
// 	res := make([]int, len(sle))
// 	copy(res, sle)
// 	return res
// }

//=========================================
// 2021-8-29 前几天面试被问到了这题，我居然完全不记得自己一个月之前做过
// 当着面试官的面，真的很难想出来，当时说了回溯算法，需要用递归或者栈来做
// 那么今天来用递归的办法做一下吧
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.7 MB, 在所有 Go 提交中击败了20.66%的用户
// 递归这个速度可以啊
//其实这个递归也是插板法的思路
// func permute(nums []int) [][]int {
// 	var part [][]int
// 	var res [][]int
// 	if len(nums) > 1 {
// 		part = permute(nums[1:])
// 	} else if len(nums) == 0 {
// 		return nil
// 	} else {
// 		res = append(res, nums)
// 		return res
// 	}
// 	l := len(part)
// 	ll := len(part[0])
// 	for i := 0; i < l; i++ {
// 		for ii := 0; ii < ll+1; ii++ {
// 			temp := []int{}
// 			if ii == 0 {
// 				temp = append(temp, nums[0])
// 				temp = append(temp, part[i]...)
// 				res = append(res, temp)
// 				continue
// 			}
// 			if ii == ll {
// 				temp = append(temp, part[i]...)
// 				temp = append(temp, nums[0])
// 				res = append(res, temp)
// 				continue
// 			}
// 			temp = append(temp, part[i][:ii]...)
// 			temp = append(temp, nums[0])
// 			temp = append(temp, part[i][ii:]...)
// 			res = append(res, temp)
// 		}
// 	}
// 	return res
// }

//===============================
// 用回溯做一遍
//2022-01-26

// func permute(nums []int) [][]int {
// 	var res [][]int
// 	var backTrack func(nums []int, numslen int, path []int)
// 	backTrack = func(nums []int, numslen int, path []int) {
// 		if len(nums) == 0 {
// 			p := make([]int, len(path))
// 			copy(p, path)
// 			res = append(res, p)
// 		}
// 		for i := 0; i < numslen; i++ {
// 			cur := nums[i]
// 			path = append(path, cur)
// 			nums = append(nums[:i], nums[i+1:]...)
// 			backTrack(nums, len(nums), path)
// 			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
// 			path = path[:len(path)-1]

// 		}
// 	}
// 	backTrack(nums, len(nums), []int{})
// 	return res

// }

//==============================
//other

var res [][]int

func permuteOther(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums, len(nums), []int{})
	return res
}
func backTrack(nums []int, numsLen int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < numsLen; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //直接使用切片
		backTrack(nums, len(nums), path)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]

	}
}
