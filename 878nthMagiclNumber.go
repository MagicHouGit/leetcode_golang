// 878. 第 N 个神奇数字
// 如果正整数可以被 A 或 B 整除，那么它是神奇的。

// 返回第 N 个神奇数字。由于答案可能非常大，返回它模 10^9 + 7 的结果。

// 示例 1：

// 输入：N = 1, A = 2, B = 3
// 输出：2
// 示例 2：

// 输入：N = 4, A = 2, B = 3
// 输出：6
// 示例 3：

// 输入：N = 5, A = 2, B = 4
// 输出：10
// 示例 4：

// 输入：N = 3, A = 6, B = 4
// 输出：8

// 提示：

// 1 <= N <= 10^9
// 2 <= A <= 40000
// 2 <= B <= 40000

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nthMagicalNumber(11, 2, 3))
	fmt.Println(nthMagicalNumber(10, 2, 3))
	// fmt.Println(nthMagicalNumber(8, 5, 7))
	// fmt.Println(nthMagicalNumber(1, 2, 3))
	// fmt.Println(nthMagicalNumber(4, 2, 3))
	// fmt.Println(nthMagicalNumber(5, 2, 4))
	// fmt.Println(nthMagicalNumber(3, 6, 4))
	// fmt.Println(nthMagicalNumber(1000000000, 40000, 40000)) //999720007
	fmt.Println(nthMagicalNumber(100000, 4, 4)) //
	// fmt.Println(nthMagicalNumber(3, 4, 4))
	fmt.Println(int(math.Pow10(15)))

}
func nthMagicalNumber(n int, a int, b int) int {
	prime := 1000000000+7
	var res int
	if a > b {
		a, b = b, a
	}
	lcm := LeastCommonMultiple(a, b)
	lcmArr := make([]int, 0)
	ma := a
	mb := b
	for ma < lcm || mb < lcm {
		if ma < mb {
			lcmArr = append(lcmArr, ma)
			ma += a
			// i = ma

		} else {
			lcmArr = append(lcmArr, mb)
			mb += b
			// i = mb
		}
	}
	zhouqi := len(lcmArr) + 1
	res = n / zhouqi * lcm
	var yushu int
	if n%zhouqi != 0 {
		yushu = lcmArr[n%zhouqi-1]
	}
	res += yushu

	return res % prime

}
func GreatestCommonDivisor(i, ii int) int {
	if ii < i {
		i, ii = ii, i
	}
	if ii%i == 0 {
		return i
	}
	return GreatestCommonDivisor(ii%i, i)

}
func LeastCommonMultiple(i, ii int) int {
	return i * ii / GreatestCommonDivisor(i, ii)
}
// func nthMagicalNumber(n int, a int, b int) int {
// 	c := a / Gcd(a, b) * b // 计算 a 和 b 的最小公倍数
// 	l, r := 1, int(1e15)
// 	for l < r {
// 		m := l + (r-l)>>1
// 		if m/a+m/b-m/c < n { // 计算 [1, m]范围内丑数的个数，若小于n，则继续扩大右区间
// 			l = m + 1
// 		} else {
// 			r = m
// 		}
// 	}
// 	return l % (1e9 + 7)
// }

// func Gcd(a, b int) int {
// 	if b == 0 {
// 		return a
// 	}
// 	return Gcd(b, a%b)
}