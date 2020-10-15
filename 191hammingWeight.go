package main

import "fmt"

func main() {
	var test uint32
	// test = 8
	test = 15
	// test = test >> 1
	// fmt.Printf("%b\n", test)
	fmt.Println(hammingWeight(test))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了49.72%的用户
func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
