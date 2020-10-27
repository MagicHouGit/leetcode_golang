// 334. 递增的三元子序列
// 给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。

// 数学表达式如下:

// 如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
// 使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
// 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。

// 示例 1:

// 输入: [1,2,3,4,5]
// 输出: true
// 示例 2:

// 输入: [5,4,3,2,1]
// 输出: false
package main

import (
	"fmt"
	"math"
)

func main() {
	// test := []int{1, 2, 3, 4, 5}
	// test := []int{1, 2, 4, 4, 5}
	// test := []int{2, 1, 5, 0, 4, 6}
	// test := []int{5, 1, 5, 5, 2, 5, 4}
	// test := []int{2, 1, 5, 0, 3}
	// test := []int{1, 1, -2, 6}
	test := []int{1, 3, 2}

	fmt.Println(increasingTriplet(test))
}

// 执行用时：4 ms, 在所有 Go 提交中击败了97.16%的用户
// 内存消耗：3 MB, 在所有 Go 提交中击败了100.00%的用户
// func increasingTriplet(nums []int) bool {
// 	l := len(nums)
// 	if l < 3 {
// 		return false
// 	}
// 	m1 := nums[0]
// 	m2 := math.MaxInt64

// 	for i := 1; i < l; i++ {
// 		if nums[i] > m2 && nums[i] > m1 {
// 			return true
// 		}
// 		if nums[i] < m2 {
// 			if nums[i] <= m1 {
// 				m1 = nums[i]
// 			} else {
// 				m2 = nums[i]
// 			}
// 		}

// 	}
// 	return false

// }
func increasingTriplet(nums []int) bool {
	l, r := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= l {
			l = nums[i]
		} else if r >= nums[i] {
			r = nums[i]
		} else {
			return true
		}
	}
	return false
}
