// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 示例:

// 输入: [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 进阶:

// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
package main

import "fmt"

func main() {
	// test := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// test := []int{-1, 2, -3, 3, 2, -1, -3}
	// test := []int{1, 2}
	// test := []int{-2, 1}
	// test := []int{-2, 1, -3, -1}
	// test := []int{2, 3, 1}
	test := []int{-2, -3, -1}
	fmt.Println(maxSubArrary(test))
}

//从首开始累加，出现最大值时，是这个最大子序和的末位，
// 从尾开始累加，出现最大值时，是这个最大子序和的首位，利用这个可以实现时间复杂度O(n)
func maxSubArrary(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {

			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
