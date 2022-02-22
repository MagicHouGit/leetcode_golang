// 89. 格雷编码
// n 位格雷码序列 是一个由 2n 个整数组成的序列，其中：
// 每个整数都在范围 [0, 2n - 1] 内（含 0 和 2n - 1）
// 第一个整数是 0
// 一个整数在序列中出现 不超过一次
// 每对 相邻 整数的二进制表示 恰好一位不同 ，且
// 第一个 和 最后一个 整数的二进制表示 恰好一位不同
// 给你一个整数 n ，返回任一有效的 n 位格雷码序列 。

// 示例 1：
// 输入：n = 2
// 输出：[0,1,3,2]
// 解释：
// [0,1,3,2] 的二进制表示是 [00,01,11,10] 。
// - 00 和 01 有一位不同
// - 01 和 11 有一位不同
// - 11 和 10 有一位不同
// - 10 和 00 有一位不同
// [0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
// - 00 和 10 有一位不同
// - 10 和 11 有一位不同
// - 11 和 01 有一位不同
// - 01 和 00 有一位不同

// 示例 2：
// 输入：n = 1
// 输出：[0,1]

// 提示：
// 1 <= n <= 16
package main

import "fmt"

func main() {
	n0 := 1
	n1 := 2
	n2 := 3
	fmt.Println(grayCode(n0))
	fmt.Println(grayCode(n1))
	fmt.Println(grayCode(n2))
	// for _, v := range grayCode(n2) {
	// 	fmt.Printf("%b\n", v)
	// }
	// t := 1 | 2
	// fmt.Println(t)

}

//backTrack
// 执行用时：8 ms, 在所有 Go 提交中击败了89.23%的用户
// 内存消耗：6.7 MB, 在所有 Go 提交中击败了62.51%的用户
//2022-02-22

func grayCode(n int) []int {
	var res []int
	var backTrack func(layer int, num int, check int)
	backTrack = func(layer int, num int, check int) {
		if layer >= n {
			res = append(res, num)
			return
		}
		injection := 1 << layer
		for i := 0; i < 2; i++ {
			if check^i == 0 {
				num = num | injection
			}
			if i == 1 {
				backTrack(layer+1, num, 0)
				num = num ^ injection
			} else {
				backTrack(layer+1, num, 1)
				num = num ^ injection
			}
		}
	}
	backTrack(0, 0, 0)
	//这两行代码是后加的,调整顺序让0在开头,
	res = append([]int{res[len(res)-1]}, res...)
	res = res[:len(res)-1]
	return res
}

//====================================
// func grayCode(n int) []int {
// 	res := make([]int, 1, 1<<n)
// 	for i := 1; i <= n; i++ {
// 		for ii := len(res) - 1; ii >= 0; ii-- {
// 			res = append(res, res[ii]|1<<(i-1))
// 		}
// 	}
// 	return res
// }

// func grayCode(n int) []int {
//     ans := make([]int, 1, 1<<n)
//     for i := 1; i <= n; i++ {
//         for j := len(ans) - 1; j >= 0; j-- {
//             ans = append(ans, ans[j]|1<<(i-1))
//         }
//     }
//     return ans
// }
