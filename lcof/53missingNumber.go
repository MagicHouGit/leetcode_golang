// 剑指 Offer 53 - II. 0～n-1中缺失的数字
// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

// 示例 1:

// 输入: [0,1,3]
// 输出: 2
// 示例 2:

// 输入: [0,1,2,3,4,5,6,7,9]
// 输出: 8

// 限制：

// 1 <= 数组长度 <= 10000
package main

import "fmt"

func main() {
	// test := []int{0, 1, 3}
	// test := []int{0, 1, 2}
	test := []int{1, 2, 3}
	fmt.Println(missingNumber(test))
}

// 执行用时：16 ms, 在所有 Go 提交中击败了95.92%的用户
// 内存消耗：6.2 MB, 在所有 Go 提交中击败了5.34%的用户
func missingNumber(nums []int) int {

	l := len(nums)
	var total int
	var count int
	for i := 0; i < l; i++ {
		total += i
		count += nums[i]
		if total != count {
			return i
		}
	}
	return l
}
