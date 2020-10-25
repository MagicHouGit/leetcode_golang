// 16. 最接近的三数之和
// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

// 示例：

// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

// 提示：

// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	test := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(test, target))
}

//也算是意料之中，这样的办法速度肯定很慢
// 执行用时：12 ms, 在所有 Go 提交中击败了19.57%的用户
// 内存消耗：2.7 MB, 在所有 Go 提交中击败了53.18%的用户
// func threeSumClosest(nums []int, target int) int {
// 	l := len(nums)
// 	res := make([]int, 3)
// 	temp := 0
// 	t := 0
// 	x := math.MaxInt64
// 	for i := 0; i < l-2; i++ {
// 		for ii := i + 1; ii < l-1; ii++ {
// 			for iii := ii + 1; iii < l; iii++ {

// 				temp = nums[i] + nums[ii] + nums[iii]
// 				if target-temp == 0 {
// 					return target
// 				}
// 				if target-temp > 0 {
// 					t = target - temp
// 				} else {
// 					t = -1 * (target - temp)
// 				}
// 				if x > t {
// 					x = t

// 					res[0], res[1], res[2] = nums[i], nums[ii], nums[iii]
// 				}

// 			}
// 		}
// 	}
// 	return res[0] + res[1] + res[2]
// }

// =======================================
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}

			if sum < target {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum > target {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
