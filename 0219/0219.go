package main

import (
	"fmt"
)

func main() {
	t0 := []int{1, 2, 3, 1}
	k0 := 3
	fmt.Println(containsNearbyDuplicate(t0, k0))
	t1 := []int{1, 0, 1, 1}
	k1 := 1
	fmt.Println(containsNearbyDuplicate(t1, k1))
	t2 := []int{1, 2, 3, 1, 2, 3}
	k2 := 2
	fmt.Println(containsNearbyDuplicate(t2, k2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	col := map[int]struct{}{}
	for i, v := range nums {
		if i > k {
			delete(col, nums[i-k-1])
		}
		if _, ok := col[v]; ok {
			return true
		}
		col[v] = struct{}{}
	}
	return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	col := make(map[int][]int, 0)
// 	for i := 0; i < len(nums); i++ {
// 		col[nums[i]] = append(col[nums[i]], i)
// 	}
// 	for _, v := range col {
// 		if len(v) > 1 {
// 			if zuixiaocha(v) <= k {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
// func zuixiaocha(nums []int) int {
// 	res := math.MaxInt64
// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i]-nums[i-1] < res {
// 			res = nums[i] - nums[i-1]
// 		}
// 	}
// 	return res
// }

//other

// func abs(n int) int {
//     if n < 0 {
//         return -n;
//     }
//     return n;
// }

// func containsNearbyDuplicate(nums []int, k int) bool {
//     var numsSize = len(nums);
//     var stack = make([]int, numsSize);
//     var top = -1;

//     for i := 0; i < numsSize; i++ {
//         for top > -1 && nums[i] < nums[stack[top]] {
//             top--;
//         }
//         if top > -1 && nums[i] == nums[stack[top]] {
//             if(abs(i - stack[top]) <= k) {
//                 return true;
//             }
//         }
//         top++;
//         stack[top] = i;
//     }

//     return false;
// }
