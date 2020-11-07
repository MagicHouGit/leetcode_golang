// 350. 两个数组的交集 II
// 给定两个数组，编写一个函数来计算它们的交集。

// 示例 1：

// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// 示例 2:

// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[4,9]

// 说明：

// 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
// 我们可以不考虑输出结果的顺序。
// 进阶：

// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
package main

import "fmt"

func main() {
	t1_nums := []int{1, 2, 2, 1, 2, 2}
	t2_nums := []int{2, 2, 2}
	intersect(t1_nums, t2_nums)
}

func intersect(nums1 []int, nums2 []int) []int {
	//取值统计,这个办法是一个范例，
	//我曾想过类似的取值统计建两个数组分别统计nums1和nums2中的情况
	//当两数组更长和数据更集中的时候会有明显的速度提升吧
	m := map[int]int{}
	var res []int
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	for _, value := range res {
		fmt.Printf("%d_\n", value)
	}
	return res
	//此法可解，但不是什么好方法，其一利用数字不为-1，其二如果两个数组长度非常长，这个算法基本就废了
	// if len(nums1) == 0 || len(nums2) == 0 {
	// 	return nil
	// }
	// reNums := []int{}

	// for i := 0; i < len(nums1); i++ {
	// 	for j := 0; j < len(nums2); j++ {
	// 		if nums1[i] == nums2[j] {
	// 			reNums = append(reNums, nums2[j])
	// 			nums2[j] = -1
	// 			break
	// 		}
	// 	}
	// }
	// for _, value := range reNums {
	// 	fmt.Println(value)
	// }
	// return reNums
}
