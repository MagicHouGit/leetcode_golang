// 344. 反转字符串
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

// 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

// 示例 1：

// 输入：["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
// 示例 2：

// 输入：["H","a","n","n","a","h"]
// 输出：["h","a","n","n","a","H"]
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
