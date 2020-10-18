package main

import "fmt"

func main() {
	var test uint32
	test = 43261596
	//964176192
	// test = 15
	// fmt.Printf()
	fmt.Println(110 & 1)
	// fmt.Println(11 & 110)
	// fmt.Println(100 & 1)
	// fmt.Println(100 & 0)
	fmt.Println(1110 | 0)
	fmt.Println(111 | 1)

	fmt.Println(reverseBits(test))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了33.25%的用户
func reverseBits(num uint32) uint32 {
	var temp uint32
	var res uint32

	for i := 0; i < 32; i++ {
		temp = num & 1
		num = num >> 1
		res = res << 1
		res = res ^ temp
	}
	return res
}
