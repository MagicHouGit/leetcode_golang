//只出现一次的数字
package main

func main() {
	ex_nums := []int{4, 1, 2, 1, 2}
	singleNumber(ex_nums)
}
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
