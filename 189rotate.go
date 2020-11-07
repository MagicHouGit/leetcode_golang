// 189. 旋转数组
// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

// 示例 1:

// 输入: [1,2,3,4,5,6,7] 和 k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右旋转 1 步: [7,1,2,3,4,5,6]
// 向右旋转 2 步: [6,7,1,2,3,4,5]
// 向右旋转 3 步: [5,6,7,1,2,3,4]
// 示例 2:

// 输入: [-1,-100,3,99] 和 k = 2
// 输出: [3,99,-1,-100]
// 解释:
// 向右旋转 1 步: [99,-1,-100,3]
// 向右旋转 2 步: [3,99,-1,-100]
// 说明:

// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 要求使用空间复杂度为 O(1) 的 原地 算法。
package main

import "fmt"

func main() {

	test_nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(test_nums, k)
}

func rotate(nums []int, k int) {
	tmp_nums := make([]int, len(nums))
	copy(tmp_nums, nums)
	if len(nums) > k {
		for i := 0; i < len(nums); i++ {
			if i < k {
				nums[i] = tmp_nums[len(nums)-k+i]
			} else {
				nums[i] = tmp_nums[i-k]
			}
		}
		for _, value := range nums {
			fmt.Printf("%d", value)
		}
	}
	if len(nums) == k {
		for _, value := range nums {
			fmt.Printf("%d", value)
		}
	}
	if len(nums) < k {
		k = k % len(nums)
		for i := 0; i < len(nums); i++ {
			if i < k {
				nums[i] = tmp_nums[len(nums)-k+i]
			} else {
				nums[i] = tmp_nums[i-k]
			}
		}
		for _, value := range nums {
			fmt.Printf("%d", value)
		}
	}

}

//这个方法len(nums)-k,会有panic
// func rotate (nums []int,k int){
// 	if len(nums)<k{
// 		for _,value := range nums{
// 			fmt.Printf("%d\n",value)
// 		}
// 	}
// 	tmp_nums := make([]int,k)
// 	for i:= 0; i<k; i++ {
// 		tmp_nums[i] = nums[len(nums)-k + i]
// 		//fmt.Printf("%d\n",tmp_nums[i])
// 	}
// 	for i:= len(nums)-1; i>=0 ;i-- {
// 		if i<k {
// 		    nums[i] = 	tmp_nums[i]
// 		}else{
//     		nums[i] = nums[i-k]
// 		}
// 	}

// 	for _,value := range nums {
// 		fmt.Printf("%d\n",value)
// 	}
// }
