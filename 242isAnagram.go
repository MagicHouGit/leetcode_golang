// 242. 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 示例 1:

// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:

// 输入: s = "rat", t = "car"
// 输出: false
// 说明:
// 你可以假设字符串只包含小写字母。

// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
package main

import (
	"fmt"
)

func main() {
	Ts := "anagram"
	Ss := "nagaran"
	fmt.Printf("%t\n", isAnagram(Ts, Ss))

}

/*
执行用时 : 8 ms, 在Valid Anagram的Go提交中击败了63.70% 的用户
内存消耗 : 3.3 MB, 在Valid Anagram的Go提交中击败了21.70% 的用户
//*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := len(s)
	tmp := make(map[rune]int, l)
	for _, v := range s {
		tmp[v]++
	}
	for _, v := range t {
		tmp[v]--
	}
	for _, v := range tmp {
		if v != 0 {
			return false
		}
	}

	return true
}

/*
	下面的方法依旧是通过之后，从leetcode摘抄的，
	但是我有一个问题，如果超出了26个字母这种方法直接就不适用了，

//*/

// func isAnagram(s string, t string) bool {
// 	sl := len(s)
// 	tl := len(t)
// 	if sl != tl {
// 		return false
// 	}
// 	if sl<1 {
// 		return true
// 	}

// 	alphabet := make([]int, 26)
// 	sb := []byte(s)
// 	tb := []byte(t)
// 	for i := 0; i < sl; i++ {
// 		alphabet[sb[i]-'a']++
// 	}
// 	for i := 0; i < tl; i++ {
// 		alphabet[tb[i]-'a']--
// 	}
// 	for i := 0; i < 26; i++ {
// 		if alphabet[i] != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
