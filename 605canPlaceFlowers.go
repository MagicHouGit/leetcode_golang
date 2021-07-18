// 605. 种花问题
// 假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

// 给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False。

// 示例 1:

// 输入: flowerbed = [1,0,0,0,1], n = 1
// 输出: True
// 示例 2:

// 输入: flowerbed = [1,0,0,0,1], n = 2
// 输出: False
// 注意:

// 数组内已种好的花不会违反种植规则。
// 输入的数组长度范围为 [1, 20000]。
// n 是非负整数，且不会超过输入数组的大小。
package main

import "fmt"

func main() {
	// testArr := []int{1, 0, 0, 0, 1, 0, 0}
	// testArr := []int{0, 0}
	// testArr := []int{0, 1}
	// testArr := []int{0}
	// testArr := []int{1}
	// testArr := []int{1, 0, 0}
	// testArr := []int{0, 0, 1}
	// testArr := []int{0, 1, 0}
	// testArr := []int{0, 0, 1, 0, 0}
	// testArr := []int{1, 0, 0, 0, 0, 1}
	testArr := []int{0, 0, 0, 0, 1}

	n := canPlaceFlowers(testArr, 2)
	fmt.Println(n)

}

// 执行用时：20 ms, 在所有 Go 提交中击败了74.11%的用户
// 内存消耗：6 MB, 在所有 Go 提交中击败了86.75%的用户
func canPlaceFlowers(flowerbed []int, n int) bool {

	max := 0
	l := len(flowerbed)
	if l == 1 {
		if flowerbed[0] == 0 {
			max = 1
		}
	}
	if l == 2 {
		if flowerbed[0]+flowerbed[1] < 1 {
			max = 1
		}
	}
	if l > 2 {

		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			max++
			flowerbed[0] = 1
		}
		for i := 1; i < l-1; i++ {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				max++
				flowerbed[i] = 1
			}
		}

		if flowerbed[l-1] == 0 && flowerbed[l-2] == 0 {
			max++
		}
	}

	fmt.Println(max)
	if max >= n {
		return true
	}
	return false
}
