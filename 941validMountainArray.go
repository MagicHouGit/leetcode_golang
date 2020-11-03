// 941. 有效的山脉数组
// 给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

// 让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

// A.length >= 3
// 在 0 < i < A.length - 1 条件下，存在 i 使得：
// A[0] < A[1] < ... A[i-1] < A[i]
// A[i] > A[i+1] > ... > A[A.length - 1]

// 示例 1：

// 输入：[2,1]
// 输出：false
// 示例 2：

// 输入：[3,5,5]
// 输出：false
// 示例 3：

// 输入：[0,3,2,1]
// 输出：true

// 提示：

// 0 <= A.length <= 10000
// 0 <= A[i] <= 10000
package main

import (
	"fmt"
)

func main() {
	// test := []int{4, 5, 2}
	// test := []int{5, 5, 2}
	// test := []int{5, 2}
	// test := []int{1, 2, 2}
	// test := []int{1, 2, 2, 3, 1}
	test := []int{1, 2, 5, 3, 3}
	fmt.Println(validMountainArray(test))
}

// 执行用时：32 ms, 在所有 Go 提交中击败了63.16%的用户
// 内存消耗：6.4 MB, 在所有 Go 提交中击败了56.58%的用户
func validMountainArray(A []int) bool {

	var mi int
	l := len(A)
	if l < 3 {
		return false
	}
	for i := 0; i < l-1; i++ {
		if A[i] < A[i+1] {
			continue
		}
		if i == 0 {
			return false
		}
		mi = i
		break
	}
	for i := mi; i < l-1; i++ {
		if A[i] > A[i+1] {
			continue
		}
		return false
	}
	return true
}

//号称是击败100%，其实时间和我一样
// func validMountainArray(A []int) bool {
// 	if len(A) < 3 {
// 		return false
// 	}

// 	tag := true

// 	for i := 0; i < len(A)-1; i++ {
// 		// 上升状态下, 前 > 后, 状态改为下降
// 		if tag && A[i] > A[i+1] {
// 			tag = false

//             if i == 0 {
// 				return false
// 			}

// 			if A[i-1] >= A[i] || A[i] <= A[i+1] {
// 				return false
// 			}
// 		}

// 		// 下降状态下, 前 < 后, 返回 false
// 		if !tag && A[i] <= A[i+1] {
// 			return false
// 		}
// 	}

// 	if tag {
// 		return false
// 	}

// 	return true
// }
