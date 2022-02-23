// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：
// 输入：nums = [0]
// 输出：[[],[0]]

// 提示：
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同

package main

import "fmt"

func main() {
	// nums0 := []int{1, 2, 3}
	// nums1 := []int{1}
	nums2 := []int{1, 2, 3, 4}

	// fmt.Println(subsets(nums0))
	// fmt.Println(subsets(nums1))
	fmt.Println(subsets(nums2))

}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了87.85%的用户
//2022-02-23
//我永远可以相信我自己会把问题想的更复杂
func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	l := len(nums)
	var backTrack func(path []int, nums []int, index int)
	backTrack = func(path []int, nums []int, index int) {
		if index == l {
			return
		}
		for i := index; i < l; i++ {
			path = append(path, nums[i])
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			// nums = append(nums[:i], nums[i+1:]...)
			backTrack(path, nums, i+1)
			// nums = append(nums[:i], append([]int{path[len(path)-1]}, nums[i:]...)...)
			path = path[:len(path)-1]
		}
	}
	backTrack([]int{}, nums, 0)
	return res

}

// func backTrack(path []int,layer int,nums []int){

// }
