// 412. Fizz Buzz
// 写一个程序，输出从 1 到 n 数字的字符串表示。

// 1. 如果 n 是3的倍数，输出“Fizz”；

// 2. 如果 n 是5的倍数，输出“Buzz”；

// 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

// 示例：

// n = 15,

// 返回:
// [
//     "1",
//     "2",
//     "Fizz",
//     "4",
//     "Buzz",
//     "Fizz",
//     "7",
//     "8",
//     "Fizz",
//     "Buzz",
//     "11",
//     "Fizz",
//     "13",
//     "14",
//     "FizzBuzz"
// ]
package main

import (
	"fmt"
	"strconv"
)

func main() {
	test := 15
	fmt.Println(fizzBuzz(test))
	// fmt.Println(1 % 3)
}

// 显示详情执行用时：4 ms, 在所有 Go 提交中击败了92.88%的用户
// 内存消耗：3.4 MB, 在所有 Go 提交中击败了84.93%的用户
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				// res= append(res, "FizzBuzz")
				res[i-1] = "FizzBuzz"
			} else {
				// res = append(res, "Fizz")
				res[i-1] = "Fizz"
			}
		} else {
			if i%5 == 0 {
				// res = append(res, "Buzz")
				res[i-1] = "Buzz"
			} else {
				// res = append(res, string(i))
				// res[i-1] = string(i)
				// res[i-1] = string(59)
				res[i-1] = strconv.Itoa(i)
			}
		}
	}
	return res

}
