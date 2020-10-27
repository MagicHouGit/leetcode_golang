// 5. 最长回文子串
// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

// 示例 1：

// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：

// 输入: "cbbd"
// 输出: "bb"

package main

import "fmt"

func main() {
	test := "gabau"
	// test := "babad"
	// test := "aa"
	// test := "ababababaa"
	// test := "a"
	// test := "cbbd"
	// test := "cdbb"
	// test := "cdb"

	// fmt.Println(isPalindrome(test))
	fmt.Println(longestPalindrome(test))
}

// 执行用时：112 ms, 在所有 Go 提交中击败了32.22%的用户
// 内存消耗：2.7 MB, 在所有 Go 提交中击败了53.54%的用户
// 选取一个字符，查找这个字符之后是否再次出现
// 如果出现那么这两个字符就可以截取出一个string，
// 这个string有可能是一个回文
// 如果这是一个唯一的字符，那么这个字符肯定不属于最大回文子串的区间
// 按照我的循环办法，这个字符不可能是回文子串的中心位
// func longestPalindrome(s string) string {

// 	l := len(s)
// 	if l <= 1 {
// 		return s
// 	}
// 	maxL := 1
// 	res := string(s[0])
// 	for i := 0; i < l-1; i++ {
// 		for ii := i + maxL; ii < l; ii++ {
// 			if s[i] == s[ii] {
// 				if isPalindrome(s[i : ii+1]) {
// 					if ii-i+1 > maxL {
// 						res = s[i : ii+1]
// 						maxL = ii - i + 1
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// //判断一个字符串是否是回文
// func isPalindrome(s string) bool {
// 	l := len(s)
// 	for i := 0; i < l/2; i++ {
// 		if s[i] == s[l-1-i] {
// 			if (l-1-i)-i <= 1 {
// 				break
// 			}
// 			continue
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

///===========================================

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	bytes := make([]byte, 2*len(s)+1)
	bytes[0] = '#'
	for i := 0; i < len(s); i++ {
		bytes[2*i+1] = s[i]
		bytes[2*i+2] = '#'
	}

	lengths := make([]int, len(bytes))
	lengths[0] = 0
	lengths[1] = 1
	maxLength := 1
	maxIndex := 1

	for i := 2; i < len(bytes); i++ {
		right := i + min(lengths[i-2], lengths[i-1])

		for right < len(bytes) && 2*i-right >= 0 {
			if bytes[right] != bytes[2*i-right] {
				break
			}
			right += 1
		}
		lengths[i] = right - i - 1

		if lengths[i] > maxLength {
			maxIndex = i
			maxLength = lengths[i]
		}
	}

	start := maxIndex - lengths[maxIndex]
	newBytes := make([]byte, lengths[maxIndex])
	if bytes[start] == '#' {
		start += 1
	}
	for i := 0; i < len(newBytes); i++ {
		newBytes[i] = bytes[start+2*i]
	}
	return string(newBytes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
