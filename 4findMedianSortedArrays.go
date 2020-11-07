// 4. 寻找两个正序数组的中位数
// 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
// 示例 1：

// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：

// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
// 示例 3：

// 输入：nums1 = [0,0], nums2 = [0,0]
// 输出：0.00000
// 示例 4：

// 输入：nums1 = [], nums2 = [1]
// 输出：1.00000
// 示例 5：

// 输入：nums1 = [2], nums2 = []
// 输出：2.00000
//

// 提示：

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//===================================================================
package main

import "fmt"

func main() {
	// test1 := []int{1, 3}
	// test1 := []int{}
	test1 := []int{1, 2}
	// test := []int{1, 2}
	// test2 := []int{2}
	// test2 := []int{}
	test2 := []int{3, 5, 6, 7, 8, 9}

	// fmt.Println(median(test))
	fmt.Println(findMedianSortedArrays(test1, test2))
}

// 执行用时：24 ms, 在所有 Go 提交中击败了26.02%的用户
// 内存消耗：5.6 MB, 在所有 Go 提交中击败了37.27%的用户
//速度上这个结果我还是挺诧异的，因为他的进阶提示是时间复杂度O(log(m+n)) 我的这个时间复杂度应该是O((m+n)/2)
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	l1 := len(nums1)
// 	l2 := len(nums2)
// 	if nums1 == nil || l1 == 0 {
// 		return median(nums2)
// 	}
// 	if nums2 == nil || l2 == 0 {
// 		return median(nums1)
// 	}
// 	lmid := (l1 + l2) / 2
// 	total := make([]int, lmid+1)
// 	for i, j := 0, 0; i+j < lmid+1; {
// 		if i == l1 {
// 			total[i+j] = nums2[j]
// 			j++
// 			continue
// 		}
// 		if j == l2 {
// 			total[i+j] = nums1[i]
// 			i++
// 			continue
// 		}
// 		if nums1[i] < nums2[j] {
// 			total[i+j] = nums1[i]
// 			// nums1 = nums1[1:]
// 			i++

// 		} else {
// 			total[i+j] = nums2[j]
// 			// nums2 = nums2[1:]
// 			j++
// 		}
// 	}
// 	if (l1+l2)%2 == 1 {

// 		return float64(total[lmid])
// 	} else {

// 		return (float64(total[lmid]) + float64(total[lmid-1])) / 2
// 	}

// }

// func median(nums []int) float64 {
// 	if nums == nil {
// 		return float64(0)
// 	}
// 	l := len(nums)
// 	if l == 0 {
// 		return float64(0)
// 	}
// 	if l == 1 {
// 		return float64(nums[0])
// 	}
// 	// total := make([]int,(l/2)+1)
// 	if l%2 == 1 {
// 		return float64(nums[l/2])
// 	} else {
// 		return (float64(nums[l/2]) + float64(nums[l/2-1])) / 2
// 	}
// }

// ==================================================================
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设nums1 短
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, len(nums1)
	k := (len(nums1) + len(nums2) + 1) / 2
	nums1Mid, nums2Mid := 0, 0

	// 二分查找
	for low <= high {
		nums1Mid = low + (high-low)>>1
		nums2Mid = k - nums1Mid

		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			low = nums1Mid + 1
		} else {
			// 找到 nums1[nums1Mid-1] <= nums2[nums2Mid] && nums1[nums1Mid] >= nums2[nums2Mid-1]
			break
		}
	}

	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}

	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}

	return float64(midLeft+midRight) / 2

}
