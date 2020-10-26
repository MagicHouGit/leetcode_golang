// 3. 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:

// 输入: "abcabcbb"
// 输出: 4
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: "bbbbb"
// 输出: 2
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

package main

import "fmt"

// TreeNode a binary tree
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func main() {
	// test := "abcabcbb"
	test := "babcbb"
	// test := "ac"
	// test := "bbbbbbb"
	// test := ""
	// test := "i"
	// test := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(test))
}

// 执行用时：4 ms, 在所有 Go 提交中击败了88.65%的用户
// 内存消耗：2.7 MB, 在所有 Go 提交中击败了68.00%的用户
// func lengthOfLongestSubstring(s string) int {

// 	l := len(s)
// 	if l <= 1 {
// 		return l
// 	}
// 	var temp []byte
// 	var maxL int
// 	temp = append(temp, s[0])
// 	for i := 1; i < l; i++ {
// 		for ii := 0; ii < len(temp); ii++ {

// 			if byte(s[i]) == temp[ii] {

// 				if len(temp) > maxL {
// 					maxL = len(temp)
// 				}
// 				temp = temp[ii+1:]
// 				break
// 			}
// 		}
// 		temp = append(temp, byte(s[i]))
// 	}
// 	if len(temp) > maxL {
// 		maxL = len(temp)
// 	}
// 	return maxL

// }

//==============================
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	freq := [128]int{}
	res := 1
	i, j := 0, 1
	freq[s[0]]++
	for i < len(s) && j < len(s) {
		for j < len(s) && freq[s[j]] == 0 {
			freq[s[j]]++
			j++
		}
		res = max(res, j-i)
		freq[s[i]]--
		i++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
