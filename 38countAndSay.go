// 38. 外观数列
// 给定一个正整数 n ，输出外观数列的第 n 项。

// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

// 你可以将其视作是由递归公式定义的数字字符串序列：

// countAndSay(1) = "1"
// countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
// 前五项如下：

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 第一项是数字 1
// 描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
// 描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
// 描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
// 描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
// 要 描述 一个数字字符串，首先要将字符串分割为 最小 数量的组，每个组都由连续的最多 相同字符 组成。然后对于每个组，先描述字符的数量，然后描述字符，形成一个描述组。要将描述转换为数字字符串，先将每组中的字符数量用数字替换，再将所有描述组连接起来。

// 例如，数字字符串 "3322251" 的描述如下图：

// 示例 1：

// 输入：n = 1
// 输出："1"
// 解释：这是一个基本样例。
// 示例 2：

// 输入：n = 4
// 输出："1211"
// 解释：
// countAndSay(1) = "1"
// countAndSay(2) = 读 "1" = 一 个 1 = "11"
// countAndSay(3) = 读 "11" = 二 个 1 = "21"
// countAndSay(4) = 读 "21" = 一 个 2 + 一 个 1 = "12" + "11" = "1211"

// 提示：

// 1 <= n <= 30
/*
执行用时 :4 ms, 在所有 Go 提交中击败了47.37%的用户
内存消耗 :6.8 MB, 在所有 Go 提交中击败了13.79%的用户
*/ //
package main

import (
	"fmt"
	"strconv"
)

func main() {
	t_n := 8
	fmt.Printf("%s\n", countAndSay(t_n))
}
func countAndSay(n int) string {
	var count int
	var tmp string = "1"
	var t string = "1"
	//数列n项就n层循环
	for i := 2; i <= n; i++ {
		count = 1
		tmp = ""
		// fmt.Printf("%d - ", len(t))
		// fmt.Printf("%s\n", t)
		//根据上一项解析出下一项。
		for j := 1; j < len(t); j++ {
			if t[j] == t[j-1] {
				count++
			} else {
				tmp += strconv.Itoa(count)
				tmp += string(t[j-1])
				count = 1
			}
		}

		tmp += strconv.Itoa(count)

		// fmt.Printf("%s\n", tmp)
		tmp += string(t[len(t)-1])
		t = tmp

	}
	return t
}
