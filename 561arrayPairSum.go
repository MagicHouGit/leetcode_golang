// 561. 数组拆分 I
// 给定长度为 2n 的整数数组 nums ，你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从 1 到 n 的 min(ai, bi) 总和最大。

// 返回该 最大总和 。

// 示例 1：

// 输入：nums = [1,4,3,2]
// 输出：4
// 解释：所有可能的分法（忽略元素顺序）为：
// 1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
// 2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
// 3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
// 所以最大总和为 4
// 示例 2：

// 输入：nums = [6,2,6,5,1,2]
// 输出：9
// 解释：最优的分法为 (2, 1), (2, 5), (6, 6). min(2, 1) + min(2, 5) + min(6, 6) = 1 + 2 + 6 = 9

// 提示：

// 1 <= n <= 104
// nums.length == 2 * n
// -104 <= nums[i] <= 104
package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 3}
	fmt.Println(arrayPairSum(arr))

}
func arrayPairSum(nums []int) int {
	res := 0
	nums = merge(nums)
	for i := 0; i < len(nums); i++ {
		if i%2 == 1 {
			res += nums[i]
		}
	}
	return res
}
func merge(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	half := length / 2
	left := merge(arr[:half])
	right := merge(arr[half:])
	return mergeSort(left, right)
}
func mergeSort(left []int, right []int) (result []int) {
	l, r := 0, 0
	for len(left) < l && len(right) < r {
		if left[l] > right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[:l]...)
	result = append(result, right[:r]...)
	return
}

//这个是速度最快的那个题解
//暂时没有看懂
func arrayPairSum(nums []int) int {
	var buckets [20001]int8
	for _, num := range nums {
		buckets[num+10000]++
	}

	sum := 0
	needAdd := true

	for num, count := range buckets {
		for count > 0 {
			if needAdd {
				sum += num - 10000
			}
			needAdd = !needAdd
			count--
		}
	}

	return sum
}
