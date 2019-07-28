/*
不知道leetcode的编译器什么毛病，tmp_byte = strs[0][i] "index out of range"
我之前的 strs[j][i] == strs[j-1][i] "index out of range"
就是各种指针越界，我的本地编译器什么问题都没有，
来气我玩下一道题去了
原来是我的问题，给for中minlen确定了范围，leetcode的编译器没有报错了
*/ //
package main

import (
	"fmt"
)

func main() {
	// tmp_strs := []string{"flower", "flow", "flight"}
	tmp_strs := []string{"l"}
	// tmp_strs := []string{"k","k"}
	fmt.Printf("%s\n", longestCommonPrefix(tmp_strs))

}
func longestCommonPrefix(strs []string) string {
	var Common_t string
	var au bool
	var tmp_byte byte
	var minlen int = 32767

	if len(strs) == 1 {
		return strs[0]
	}
	for _, v := range strs {
		if l := len(v); l < minlen {
			minlen = l
		}
	}
	for i := 0; i < minlen; i++ {

		for j := 1; j < len(strs); j++ {
			tmp_byte = strs[0][i]
			if strs[j][i] == tmp_byte {

				au = true
			} else {
				au = false
				break
			}
		}
		if au == true {
			Common_t += string(strs[0][i])
		} else {
			break
		}

	}
	return Common_t

	// if len(strs) == 0 {
	// 	return ""
	// }

	// minlen := 32767

	// for _, s := range strs {
	// 	if l := len(s); l < minlen {
	// 		minlen = l
	// 	}
	// }
	// i := 0
	// for ; i < minlen; i++ {
	// 	n := strs[0][i]
	// 	j := 1

	// 	for ; j < len(strs); j++ {
	// 		if n != strs[j][i] {
	// 			break
	// 		}
	// 		n &= strs[j][i]
	// 	}

	// 	if j < len(strs) {
	// 		break
	// 	}
	// }

	// if i <= 0 {
	// 	return ""
	// }

	// return strs[0][:i]
}
