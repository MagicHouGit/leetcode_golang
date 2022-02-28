// 47. 全排列 II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
// 示例 1：
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]]
// 示例 2：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// 提示：
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10

package main

import "fmt"

func main() {
	nums0 := []int{1, 2, 2, 3}
	fmt.Println(permuteUnique(nums0))
}

//backTrack+prune
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.9 MB, 在所有 Go 提交中击败了22.66%的用户
//2022-02-28
//剪枝的思路很准
func permuteUnique(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	var backTrack func(nums []int, path []int, index int, lenNums int)
	backTrack = func(nums []int, path []int, index int, lenNums int) {
		if l == len(path) {
			temp := make([]int, l)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		//数 - 位置
		h := make(map[int]int)
		for i := 0; i < lenNums; i++ {
			if _, ok := h[nums[i]]; ok {
				continue
			}
			pick := nums[i]
			path = append(path, pick)
			nums = append(nums[:i], nums[i+1:]...)
			backTrack(nums, path, i+1, lenNums-1)
			nums = append(nums[:i], append([]int{pick}, nums[i:]...)...)
			path = path[:len(path)-1]
			h[nums[i]] = len(path)
		}
	}
	backTrack(nums, []int{}, 0, l)
	return res
}
