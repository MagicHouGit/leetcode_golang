// 125. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:

// 输入: "race a car"
// 输出: false
package main

import (
	"fmt"
	"strings"
)

func main() {
	// TmpS := "A man, a plan, a canal: Panama"
	TmpS := "0P"
	fmt.Printf("%t\n", isPalindrome(TmpS))
}

/*
执行用时 : 4 ms, 在Valid Palindrome的Go提交中击败了97.76% 的用户
内存消耗 : 4.6 MB, 在Valid Palindrome的Go提交中击败了24.33% 的用户
//*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	tmp := []rune{}
	for _, v := range s {
		if v >= 97 && v <= 122 {
			tmp = append(tmp, v)
		}
		if v >= 48 && v <= 57 {
			tmp = append(tmp, v)
		}
	}
	l := len(tmp)
	//进行首尾比较
	for i := 0; i < (l+1)/2; i++ {
		if tmp[i] != tmp[l-i-1] {
			return false
		}
	}
	// for i, v := range tmp {
	// 	fmt.Printf("%d-%c\n", i, v)
	// }
	return true

}

/*
借鉴一下leetcode上的其他解法
//*/
func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		sc, ec := s[start], s[end]
		if sc >= 'A' && sc <= 'Z' {
			sc += 32
		}
		if ec >= 'A' && ec <= 'Z' {
			ec += 32
		}

		if !((sc >= '0' && sc <= '9') || (sc >= 'a' && sc <= 'z')) {
			start++
			continue
		}
		if !((ec >= '0' && ec <= '9') || (ec >= 'a' && sc <= 'z')) {
			end--
			continue
		}
		if sc != ec {
			return false
		}
		start++
		end--
	}
	return true
}
