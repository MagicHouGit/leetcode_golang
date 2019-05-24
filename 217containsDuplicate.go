//存在重复元素
package main

import "fmt"

func main() {
	test_nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(test_nums))
}
func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}

		}
	}
	return false
}

// 此方法来自leetcode
// func containsDuplicate(nums []int) bool {
// 	m := make(map[int]bool)
// 	for _, i := range nums {
// 		if _, ok := m[i]; ok {
// 			return true
// 		}
// 		m[i] = true
// 	}
//  	return false
// }
