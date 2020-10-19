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
