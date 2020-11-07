// 461. 汉明距离
// 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。

// 给出两个整数 x 和 y，计算它们之间的汉明距离。

// 注意：
// 0 ≤ x, y < 231.

// 示例:

// 输入: x = 1, y = 4

// 输出: 2

// 解释:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑

// 上面的箭头指出了对应二进制位不同的位置。
package main

import "fmt"

func main() {
	// a := 1
	// b := 4
	b := 15
	a := 1
	fmt.Println(hammingDistance(a, b))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了49.24%的用户
func hammingDistance(x int, y int) int {
	c := x ^ y
	n := 0
	for {
		if c&1 == 1 {
			n++
		}
		c = c >> 1
		if c == 0 {
			break
		}
	}
	return n

}
