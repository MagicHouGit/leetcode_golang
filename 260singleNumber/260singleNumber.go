// 260. 只出现一次的数字 III
// 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
// 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
// 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

package singleNumber

// 执行用时：8 ms, 在所有 Go 提交中击败了27.51%的用户
// 内存消耗：4.4 MB, 在所有 Go 提交中击败了18.78%的用户
// 通过测试用例：32 / 32

//看了一眼,讨论区,第一个就是用异或的,我就知道我用map就败了
func singleNumber(nums []int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := count[nums[i]]; ok {
			delete(count, nums[i])
		} else {
			count[nums[i]] = 1
		}
	}
	var res []int
	for k := range count {
		res = append(res, k)
	}
	return res
}

// =============================

// 执行用时：4 ms, 在所有 Go 提交中击败了97.41%的用户
// 内存消耗：3.9 MB, 在所有 Go 提交中击败了100.00%的用户
func singleNumber1(nums []int) []int {
	x := 0
	for _, v := range nums {
		x ^= v
	}
	tar := 1
	for tar&x == 0 {
		tar <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		if v&tar != 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}
