// 326. 3的幂
// 给定一个整数，写一个函数来判断它是否是 3 的幂次方。

// 示例 1:

// 输入: 27
// 输出: true
// 示例 2:

// 输入: 0
// 输出: false
// 示例 3:

// 输入: 9
// 输出: true
// 示例 4:

// 输入: 45
// 输出: false
// 进阶：
// 你能不使用循环或者递归来完成本题吗？
package main

import "fmt"

func main() {
	test := 6951
	fmt.Println(isPowerOfThree(test))
}

func isPowerOfThree(n int) bool {

	if n == 0 {
		return false
	}
	p := 1
	for p < n {
		p = p * 3
	}
	if p == n {
		return true
	}
	return false
}

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		}
		if n%3 != 0 {
			return false
		}
		n = n / 3
		if n <= 0 {
			return false
		}
	}
	return false
}
