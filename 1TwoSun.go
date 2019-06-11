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
