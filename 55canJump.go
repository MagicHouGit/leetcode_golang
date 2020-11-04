// 55. 跳跃游戏
// 给定一个非负整数数组，你最初位于数组的第一个位置。

// 数组中的每个元素代表你在该位置可以跳跃的最大长度。

// 判断你是否能够到达最后一个位置。

// 示例 1:

// 输入: [2,3,1,1,4]
// 输出: true
// 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
// 示例 2:

// 输入: [3,2,1,0,4]
// 输出: false
// 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
package main

import "fmt"

func main() {
	// test := []int{2, 3, 1, 1, 4}
	// test := []int{2}
	// test := []int{0, 0}
	test := []int{1, 0, 0}
	// test := []int{0, 1, 1, 2, 0}
	// test := []int{3, 2, 1, 0, 4}
	// test := []int{1, 1, 2, 0, 0}
	fmt.Println(canJump(test))
}

// 执行用时：8 ms, 在所有 Go 提交中击败了91.44%的用户
// 内存消耗：4 MB, 在所有 Go 提交中击败了100.00%的用户
func canJump(nums []int) bool {
	l := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < l {
			l++
		} else {
			l = 1
		}
	}
	if l == 1 {
		return true
	}

	return false
}
