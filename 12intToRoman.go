// 12. 整数转罗马数字
// 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

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
// 给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

// 示例 1:

// 输入: 3
// 输出: "III"
// 示例 2:

// 输入: 4
// 输出: "IV"
// 示例 3:

// 输入: 9
// 输出: "IX"
// 示例 4:

// 输入: 58
// 输出: "LVIII"
// 解释: L = 50, V = 5, III = 3.
// 示例 5:

// 输入: 1994
// 输出: "MCMXCIV"
// 解释: M = 1000, CM = 900, XC = 90, IV = 4.
package main

import "fmt"

func main() {
	// test := 200
	// test := 3
	// test := 4
	// test := 9
	// test := 58
	test := 1994
	fmt.Println(intToRoman(test))
}

// 执行用时：8 ms, 在所有 Go 提交中击败了84.04%的用户
// 内存消耗：3.4 MB, 在所有 Go 提交中击败了45.20%的用户

// func intToRoman(num int) string {

// 	arr := [4]int{}

// 	arr[0] = num / 1000
// 	arr[1] = (num - arr[0]*1000) / 100
// 	arr[2] = (num - arr[0]*1000 - arr[1]*100) / 10
// 	arr[3] = num - arr[0]*1000 - arr[1]*100 - arr[2]*10
// 	var a, b, c string
// 	var res string
// 	for i := 0; i < 4; i++ {
// 		a, b, c = re(i)
// 		// re = append(re,nm(a,b,c,arr[i]))
// 		res = res + nm(a, b, c, arr[i])
// 	}
// 	return res

// }
// func nm(a string, b string, c string, n int) string {
// 	switch n {
// 	case 1:
// 		return a
// 	case 2:
// 		return a + a
// 	case 3:
// 		return a + a + a
// 	case 4:
// 		return a + b
// 	case 5:
// 		return b
// 	case 6:
// 		return b + a
// 	case 7:
// 		return b + a + a
// 	case 8:
// 		return b + a + a + a
// 	case 9:
// 		return a + c
// 	}
// 	return ""
// }

// func re(n int) (string, string, string) {
// 	switch n {
// 	case 0:
// 		return "M", "", ""
// 	case 1:
// 		return "C", "D", "M"
// 	case 2:
// 		return "X", "L", "C"
// 	case 3:
// 		return "I", "V", "X"
// 	}
// 	return "", "", ""
// }
func intToRoman(num int) string {
	var arr []int
	var div, rem int
	var ret string
	const_5 := []string{"V", "L", "D"}
	const_1 := []string{"I", "X", "C", "M"}

	getS := func(num, digit int) string {
		if num == 0 {
			return ""
		}
		if num == 9 {
			return const_1[digit] + const_1[digit+1]
		}
		if num == 4 {
			return const_1[digit] + const_5[digit]
		}
		if num == 5 {
			return const_5[digit]
		}
		if num <= 3 {
			tmp := ""
			for i := 1; i <= num; i++ {
				tmp += const_1[digit]
			}
			return tmp
		}

		tmp := const_5[digit]
		for i := 6; i <= num; i++ {
			tmp += const_1[digit]
		}
		return tmp
	}

	for {
		if num == 0 {
			break
		}
		div, rem = num/10, num%10
		arr = append(arr, rem)
		num = div
	}

	for i := len(arr) - 1; i >= 0; i-- {
		ret = ret + getS(arr[i], i)
	}
	return ret
}
