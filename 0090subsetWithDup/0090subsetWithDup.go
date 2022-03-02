// 90. 子集 II
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 示例 1：
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
// 示例 2：
// 输入：nums = [0]
// 输出：[[],[0]]
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums0 := []int{1, 2, 2}
	nums1 := []int{0}

	fmt.Println(subsetsWithDup(nums0))
	fmt.Println(subsetsWithDup(nums1))
}

//backTrack+prune
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了24.88%的用户
//2022-03-01
//修剪同一支,修剪同一层,的区分
// func subsetsWithDup(nums []int) [][]int {
// 	var res [][]int
// 	res = append(res, []int{})
// 	h := make(map[int]bool)
// 	var backTrack func(path []int, index int)
// 	backTrack = func(path []int, index int) {

// 		for i := index; i < len(nums); i++ {
// 			if i > 0 && nums[i-1] == nums[i] && h[i-1] == false {
// 				continue
// 			}
// 			path = append(path, nums[i])
// 			temp := make([]int, len(path))
// 			copy(temp, path)
// 			res = append(res, temp)
// 			h[i] = true
// 			backTrack(path, i+1)
// 			path = path[:len(path)-1]
// 			h[i] = false
// 		}
// 	}
// 	sort.Ints(nums)
// 	backTrack([]int{}, 0)
// 	return res
// }

//=======================================
//other
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了99.76%的用户
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	r := make([]int, 0)
	var dp func(int, int)
	sort.Ints(nums)
	dp = func(i, last int) {
		if i == len(nums) {
			t := make([]int, len(r))
			copy(t, r)
			res = append(res, t)
			return
		}
		if nums[i] >= last {
			r = append(r, nums[i])
			dp(i+1, nums[i])
			r = r[:len(r)-1]
		}
		if nums[i] != last {
			dp(i+1, last)
		}
	}
	dp(0, -11)
	return res
}
