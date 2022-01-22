package main

import "fmt"

func main() {
	n0 := 1
	n1 := 6
	n2 := 9
	fmt.Println(lastRemaining(n0))
	fmt.Println(lastRemaining(n1))
	fmt.Println(lastRemaining(n2))
	// fmt.Println(lastRemaining_other(n))

}
func lastRemaining(n int) int {
	head := 1
	cnt := n
	step := 1
	k := 0
	for cnt != 1 {
		if k%2 == 0 {
			head = head + step
		} else {
			if cnt%2 != 0 {
				head = head + step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return head
}

//==================
// other idea
// func lastRemaining_other(n int) int {
// 	a1 := 1
// 	k, cnt, step := 0, n, 1
// 	for cnt > 1 {
// 		if k%2 == 0 { // 正向
// 			a1 += step
// 		} else { // 反向
// 			if cnt%2 == 1 {
// 				a1 += step
// 			}
// 		}
// 		k++
// 		cnt >>= 1
// 		step <<= 1
// 	}
// 	return a1
// }
