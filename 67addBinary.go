// 67. 二进制求和

// 给定两个二进制字符串，返回他们的和（用二进制表示）。

// 输入为非空字符串且只包含数字 1 和 0。

// 示例 1:

// 输入: a = "11", b = "1"
// 输出: "100"
// 示例 2:

// 输入: a = "1010", b = "1011"
// 输出: "10101"

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-binary
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
)

func main() {
	a := "11"
	b := "11"
	// addBinary(a, b)
	// fmt.Println(a[2])
	// a = append(a, "1")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(addBinary(a, b))
	// fmt.Println(len(a))
}

// 执行用时 :4 ms, 在所有 Go 提交中击败了16.75%的用户
// 内存消耗 :2.6 MB, 在所有 Go 提交中击败了26.09%的用户
func addBinary(a string, b string) string {
	var lMax, lMin int
	var Max, Min string
	if len(a) > len(b) {
		lMax = len(a)
		Max = a
		lMin = len(b)
		Min = b
	} else {
		lMax = len(b)
		Max = b
		lMin = len(a)
		Min = a
	}
	var sup string
	for i := 0; i < lMax-lMin; i++ {
		sup += "0"
	}
	Min = sup + Min
	var wei string = "0"
	var res string
	for i := lMax - 1; i >= 0; i-- {
		if Max[i] == '1' && Min[i] == '1' && wei == "1" {
			wei = "1"
			res = "1" + res
			continue
		}
		if (Max[i] == '1' && Min[i] == '1' && wei == "0") || (wei == "1" && (Max[i] == '1' || Min[i] == '1')) {
			wei = "1"
			res = "0" + res
			continue
		}
		if Max[i] == '0' && Min[i] == '0' && wei == "0" {
			wei = "0"
			res = "0" + res
			continue
		}
		if wei == "1" || Max[i] == '1' || Min[i] == '1' {
			wei = "0"
			res = "1" + res
			continue
		}

	}
	if wei == "1" {
		res = "1" + res
	}
	return res
}
