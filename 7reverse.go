//7整数反转
package main

import "fmt"

func main() {
	tempX := -2147483412
	fmt.Printf("%d\n", reverse(tempX))
}

// // 执行用时 : 0 ms, 在Reverse Integer的Go提交中击败了100.00% 的用户
// // 内存消耗 : 2.2 MB, 在Reverse Integer的Go提交中击败了68.95% 的用户
// func reverse(x int) int {
// 	// long := len(x)
// 	num := make([]int, 10)
// 	tX := x
// 	var rtX int = 0
// 	numBit := 0
// 	for tX != 0 {
// 		tX /= 10
// 		numBit++
// 	}
// 	// zeroBit := 0
// 	tX = x
// 	//把每一位放入数组
// 	for i := 0; i < numBit; i++ {
// 		num[i] = tX % 10
// 		tX = tX / 10
// 	}

// 	for i := 0; i < numBit; i++ {
// 		if i == 9 {
// 			if rtX > 214748364 || rtX < -214748364 {
// 				return 0
// 			}
// 			if rtX == 214748364 && num[9] > 7 {
// 				return 0
// 			}
// 			if rtX == -214748364 && num[9] < -8 {
// 				return 0
// 			}
// 		}

// 		rtX *= 10
// 		rtX = rtX + num[i]
// 	}
// 	return rtX
// }

//-----------------------------------------------------------------------------

//果然还是要抄抄别人的代码
//别人的代码确实有我没有想到的点，但是我认为这段代码有问题，如果反转之后超过了int的存储最大值，
//那么Maxint大的数，是如何被存储的呢，应该已经溢出了
func reverse(x int) int {
	var left = x
	var newint = 0
	var Maxint = 1<<31 - 1
	var Minint = -1 << 31
	for {
		if left == 0 {
			break
		}

		newint = newint*10 + left%10
		left = left / 10
		if newint < Minint || newint > Maxint {
			return 0
		}
	}
	return newint
}
