//15三数之和

// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

// 示例：

// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

package main

import "fmt"

func main() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))

}

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
func threeSum(nums []int) [][]int {
	nums = mergeSort(nums)
	fmt.Println(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		k := len(nums) - 1
		j := i + 1
		if i == 0 || nums[i] != nums[i-1] {

			for j < k {
				s := nums[i] + nums[j] + nums[k]
				if s == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					fmt.Printf("%d,%d,%d,\n", i, j, k)
					res = append(res, temp)
					j++
					k--
					for j < k && nums[j] == nums[j-1] {
						j++
					}
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				}
				if s > 0 {
					j++
				}
				if s < 0 {
					k--
				}
			}
		}
	}
	return res
}
