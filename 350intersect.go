//两个数组的交集2
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
