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
