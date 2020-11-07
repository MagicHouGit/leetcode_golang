// 1. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
package main

import "fmt"

func main() {
	t_nums := []int{2, 11, 5, 6}
	t_target := 16
	reNums := TwoSum(t_nums, t_target)
	for _, v := range reNums {
		fmt.Printf("%d\n", v)
	}
}

// //执行用时 : 92 ms, 在Two Sum的Go提交中击败了14.22% 的用户
// //内存消耗 : 3 MB, 在Two Sum的Go提交中击败了78.72% 的用户
// //这个简直就是耻辱，
// func TwoSum(nums []int, target int) []int {
// 	size := len(nums)
// 	rNums := []int{0, 1}
// 	pl := 0
// 	for i := 0; i < size-1; i++ {
// 		for j := i + 1; j < size; j++ {
// 			pl = nums[i] + nums[j]
// 			if pl == target {
// 				rNums[0], rNums[1] = i, j
// 				return rNums
// 			}
// 		}
// 	}
// 	return rNums
// }

//这个来自leecode的方法
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[n]; ok {
			return []int{j, i}
		}
		m[target-n] = i
	}
	return []int{}
}
