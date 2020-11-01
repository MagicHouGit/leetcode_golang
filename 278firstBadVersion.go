package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(2))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了68.16%的用户
func firstBadVersion(n int) int {

	return er(1, n)
}
func er(u, t int) int {
	if t-u <= 1 {
		if isBadVersion(u) == true {
			return u
		} else {
			return t
		}
	}
	mid := (u + t) / 2
	if isBadVersion(mid) == true {
		return er(u, mid)
	} else {
		return er(mid, t)
	}
}

func isBadVersion(version int) bool {
	if version > 1 {
		return true
	} else {
		return false
	}
}
