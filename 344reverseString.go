//反转字符串
package main

import "fmt"

func main() {
	temp_s := []byte{'h', 'o'}
	reverseString(temp_s)
	for _, v := range temp_s {
		fmt.Printf("%x,", v)
	}

	fmt.Println("\n")
}

//执行用时 : 1840 ms, 在Reverse String的Go提交中击败了80.10% 的用户
//内存消耗 : 19.2 MB, 在Reverse String的Go提交中击败了49.23% 的用户
//我这里用的string类型，但是leetcode上byte，我电脑上的编译器byte类型报错
//后来发现是我的问题，单引号，双引号区别
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
