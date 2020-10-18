package main

import "fmt"

func main() {
	// a := 1
	// b := 4
	b := 15
	a := 1
	fmt.Println(hammingDistance(a, b))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了49.24%的用户
func hammingDistance(x int, y int) int {
	c := x ^ y
	n := 0
	for {
		if c&1 == 1 {
			n++
		}
		c = c >> 1
		if c == 0 {
			break
		}
	}
	return n

}
