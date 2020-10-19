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
