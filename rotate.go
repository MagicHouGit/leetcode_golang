//旋转数组

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
