// 13. 罗马数字转整数
// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
// 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

// 示例 1:

// 输入: "III"
// 输出: 3
// 示例 2:

// 输入: "IV"
// 输出: 4
// 示例 3:

// 输入: "IX"
// 输出: 9
// 示例 4:

// 输入: "LVIII"
// 输出: 58
// 解释: L = 50, V= 5, III = 3.
// 示例 5:

// 输入: "MCMXCIV"
// 输出: 1994
// 解释: M = 1000, CM = 900, XC = 90, IV = 4.

// 提示：

// 题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
// IC 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
// 关于罗马数字的详尽书写规则，可以参考 罗马数字 - Mathematics 。

package main

import "fmt"

func main() {
	// test := "III"
	// test1 := "01234"
	// test := "IV"
	// test := "V"
	// test := "LVIII"
	// test := "LIV"
	test := "MCMXCIV"
	// fmt.Println(test1[0:2])
	// fmt.Println(test[2])
	// fmt.Printf("%s\n", test[1])
	// fmt.Printf("%c\n", test[1])
	// fmt.Printf("%p\n", test[1])
	fmt.Println(romanToInt(test))
}

// 执行用时：16 ms, 在所有 Go 提交中击败了23.32%的用户
// 内存消耗：5 MB, 在所有 Go 提交中击败了5.03%的用户
// func romanToInt(s string) int {
// 	if s == "" {
// 		return 0
// 	}
// 	index := make(map[string]int)
// 	index["I"] = 1
// 	index["V"] = 5
// 	index["X"] = 10
// 	index["L"] = 50
// 	index["C"] = 100
// 	index["D"] = 500
// 	index["M"] = 1000
// 	index["IV"] = 4
// 	index["IX"] = 9
// 	index["XL"] = 40
// 	index["XC"] = 90
// 	index["CD"] = 400
// 	index["CM"] = 900
// 	var res int
// 	for i := 0; i < len(s); {
// 		if len(s) < 2 {
// 			break
// 		}
// 		_, ok := index[s[0:2]]
// 		if ok {
// 			res += index[string(s[0:2])]
// 			s = s[2:]
// 		} else {
// 			res += index[string(s[0])]
// 			s = s[1:]
// 		}
// 	}
// 	if len(s) == 1 {
// 		res += index[string(s[0])]
// 	}

// 	return res
// }

//----------------------
func romanToInt(s string) (r int) {

	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := judge(s[i])
		if cur == -1 {
			return -1
		} else {
			if cur >= pre {
				r += cur
			} else {
				r -= cur
			}
			pre = cur
		}

	}
	return r
}

func judge(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return -1
	}
}
