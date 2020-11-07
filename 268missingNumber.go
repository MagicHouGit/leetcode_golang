// 268. 丢失的数字
// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

// 进阶：

// 你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?

// 示例 1：

// 输入：nums = [3,0,1]
// 输出：2
// 解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
// 示例 2：

// 输入：nums = [0,1]
// 输出：2
// 解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
// 示例 3：

// 输入：nums = [9,6,4,2,3,5,7,0,1]
// 输出：8
// 解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
// 示例 4：

// 输入：nums = [0]
// 输出：1
// 解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。

// 提示：

// n == nums.length
// 1 <= n <= 104
// 0 <= nums[i] <= n
// nums 中的所有数字都 独一无二
package main

import "fmt"

func main() {
	// test := []int{0, 1, 3}
	// test := []int{0, 1, 2}
	// test := []int{1, 2, 3}
	// test := []int{0, 1, 2, 3, 4, 5, 6}
	test := []int{0, 5, 2, 6, 4, 1, 3}
	fmt.Println(missingNumber(test))
	t := []int{0, 1, 2, 3, 4}
	p := []int{3, 1, 0, 2, 4}
	var x int
	for i := 0; i < 5; i++ {
		x = x ^ t[i]
		x = x ^ p[i]
		fmt.Println(x)
	}
	// fmt.Println(7 ^ 1)
}

// 执行用时：20 ms, 在所有 Go 提交中击败了76.75%的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了42.35%的用户
// func missingNumber(nums []int) int {

// 	l := len(nums)
// 	var total int
// 	var count int
// 	for i := 0; i < l; i++ {
// 		total += i
// 		count += nums[i]
// 	}
// 	if total == count {
// 		return l
// 	}
// 	total += l

// 	return total - count
// }

//=========================================

func missingNumber(nums []int) int {
	missing := len(nums)
	var temp int

	for i := 0; i < len(nums); i++ {
		// missing ^= i ^ nums[i]
		temp = i ^ nums[i]
		missing = missing ^ temp
	}
	return missing

}
