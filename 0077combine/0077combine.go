package main

import "fmt"

func main() {
	// n0 := 4
	// k0 := 2
	// fmt.Println(combine(n0, k0))
	fmt.Println(combine(1, 1))
}

//backTrack
// 执行用时：20 ms, 在所有 Go 提交中击败了14.95%的用户
// 内存消耗：9.8 MB, 在所有 Go 提交中击败了9.45%的用户
//第一次独立写出回溯,这个速度也是意料之中,
var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	heap := []int{}
	for i := 0; i < n; i++ {
		heap = append(heap, i+1)
	}
	backTrack(heap, k, []int{}, []int{})
	return res
}
func backTrack(heap []int, k int, used []int, subset []int) {
	if len(subset) == k {
		temp := make([]int, k)
		copy(temp, subset)
		res = append(res, temp)
	} else {
		n := len(heap)
		for i := 0; i < n; i++ {
			used = append(used, heap[i])
			subset = append(subset, heap[i])
			backTrack(heap[i+1:], k, used, subset)
			used = used[:len(used)-1]
			subset = subset[:len(subset)-1]
		}
	}
}
