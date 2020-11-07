// 387. 字符串中的第一个唯一字符
// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

// 示例：

// s = "leetcode"
// 返回 0

// s = "loveleetcode"
// 返回 2

// 提示：你可以假定该字符串只包含小写字母。
package main

import (
	"fmt"
)

func main() {
	s := "eetcodeleetcode"
	fmt.Printf("%d\n", firstUniqChar(s))
}
func firstUniqChar(s string) int {
	// hep := []stringo
	//map的方式有些不太适用，原因：找出第一个
	// hep := make(map[string]int, len(s))
	// for _, v := range s {
	// 	hep[string(v)]++
	// }
	// for i, v := range hep {
	// 	fmt.Println(i, v)
	// 	if v == 1
	// }
	//--------------------------------------------------
	//想来想去感觉，如果是找出第一个不重复的字符，双循环的方式也许是最好的
	//执行用时 : 76 ms, 在First Unique Character in a String的Go提交中击败了22.44% 的用户
	//内存消耗 : 6.9 MB, 在First Unique Character in a String的Go提交中击败了5.26% 的用户
	//这个速度是我不能接受的，必须重写。
	// l := len(s)
	// UcBool := false
	// tmp := make([]string, l)
	// c := 0
	// for _, v := range s {
	// 	tmp[c] = string(v)
	// 	c++
	// }
	// for i := 0; i < l; i++ {
	// 	for j := 0; j < l; j++ {
	// 		if i == j {
	// 			continue
	// 		}
	// 		if tmp[i] == tmp[j] {
	// 			UcBool = true
	// 			break
	// 		}

	// 	}
	// 	if UcBool == false {
	// 		return i
	// 	}
	// 	UcBool = false
	// }
	// return -1
	//---------------------------------------------------------
	/*
		还是要重新利用map的方法了，
		执行用时 : 60 ms, 在First Unique Character in a String的Go提交中击败了42.90% 的用户
		内存消耗 : 8.6 MB, 在First Unique Character in a String的Go提交中击败了5.26% 的用户
		不行啊，这个速度我还是不能接受
	//*/
	// l := len(s)
	// hep := make(map[string]int, l)
	// tmp := make([]string, l)
	// c := 0
	// for _, v := range s {
	// 	hep[string(v)]++
	// 	tmp[c] = string(v)
	// 	c++
	// }
	// for i := 0; i < l; i++ {
	// 	if hep[tmp[i]] == 1 {
	// 		return i
	// 	}
	// }
	// return -1
	m := make([]int, 26)
	for _, char := range s {
		m[char-'a']++
	}
	for idx, char := range s {
		if m[char-'a'] == 1 {
			return idx
		}
	}
	return -1

}
