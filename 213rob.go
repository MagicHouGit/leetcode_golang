// 213. 打家劫舍 II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。

// 示例 1：

// 输入：nums = [2,3,2]
// 输出：3
// 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
// 示例 2：

// 输入：nums = [1,2,3,1]
// 输出：4
// 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//      偷窃到的最高金额 = 1 + 3 = 4 。
// 示例 3：

// 输入：nums = [0]
// 输出：0

// 提示：

// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
package main

import "fmt"

func main() {
	// test := []int{2, 3, 1}
	// test := []int{2, 7, 9, 3, 1, 4}
	// test := []int{}
	test := []int{0, 0}
	// test := []int{2, 3, 2}
	// test := []int{1, 2, 1, 1}
	fmt.Println(rob(test))

}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了16.39%的用户
func rob(nums []int) int {

	if nums == nil {
		return 0
	}
	l := len(nums)
	// 我都蒙了，nil居然不是长度==0
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	res := make([]int, l)
	res[0] = nums[0]
	res[1] = max(nums[0], nums[1])

	if l == 2 {
		return res[1]
	}
	if l == 3 {
		return max(max(nums[0], nums[1]), max(nums[1], nums[2]))
	}
	res1 := make([]int, l-1)
	res1[0] = nums[1]
	res1[1] = max(nums[1], nums[2])
	for i := 2; i < l-1; i++ {
		// m, n := i-1, i+1

		res[i] = max(res[i-2]+nums[i], res[i-1])
		res1[i] = max(res1[i-2]+nums[i+1], res1[i-1])
	}
	return max(res[l-1-1], res1[l-1-1])

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
