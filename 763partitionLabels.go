// 763. 划分字母区间
// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

// 示例：

// 输入：S = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca", "defegde", "hijhklij"。
// 每个字母最多出现在一个片段中。
// 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

// 提示：

// S的长度在[1, 500]之间。
// S只包含小写字母 'a' 到 'z' 。
package main

import "fmt"

func main() {

	// S := "ababcbacadefegdehijhklij"
	S := "ab"
	// S := "abc"
	// S := "abac"
	// S := "a"
	// S := ""
	// S := "aa"
	// S := "abab"
	// S := "caedbdedda"
	// to := make(map[string][2]int, 26)
	// fmt.Println(to)
	fmt.Println(len(S))
	fmt.Println(partitionLabels(S))
}

//调试花了不少时间，a,ab,abc,abac,这个几种情况不能兼顾
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了20.00%的用户

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return nil
	}
	//记录不重复的字母
	var repeat []byte
	//和repeat的位置对应，这个字母的起始位置和结束位置，
	var to [][]int
	var res []int

	//
	repeat = append(repeat, S[0])
	to = append(to, []int{0, 0})
	for i, v := range S {
		for k := 0; k < len(repeat); k++ {
			if byte(v) == repeat[k] {
				to[k][1] = i
				break
			}
			if k == len(repeat)-1 {
				repeat = append(repeat, byte(v))
				to = append(to, []int{i, i})
				break
			}
		}
	}
	fmt.Println(repeat)
	fmt.Println(to)

	remain := len(S)
	start, end := to[0][0], to[0][1]
	for i := 1; i < len(repeat); i++ {

		// if to[i][1] == 0 {
		// 	// if to[i][0] == 0 {
		// 	// 	res = append(res, 1)
		// 	// 	start, end = 1, 1
		// 	// 	continue
		// 	// }
		// 	to[i][1] = to[i][0]
		// }
		// if start == end && i < len(repeat)-1 {
		// 	res = append(res, 1)
		// 	start = to[i+1][0]
		// 	end = to[i+1][1]
		// 	continue
		// }
		//这个字符的起始是否大于之前字符的结束，
		//大于就代表一个段落的结束，从当前字符开启一个新的段落
		if to[i][0] > end {
			res = append(res, end-start+1)
			remain -= end - start + 1
			start = to[i][0]
			end = to[i][1]
			continue
		}
		if to[i][1] > end {
			end = to[i][1]
		}
		// if i == len(repeat)-1 {
		// 	res = append(res, end-start+1)
		// }
	}

	if remain > 0 {
		res = append(res, remain)
	}
	return res
}
