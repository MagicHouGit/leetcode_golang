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
