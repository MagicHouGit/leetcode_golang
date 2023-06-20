package main

import "fmt"

func main() {
	arr1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(arr1))
	for _, v := range arr1 {
		if v*2 == v<<1 {
			fmt.Println("tru")
		} else {
			fmt.Println("false")
		}
	}
}

// func findDuplicates(nums []int) []int {
// 	// var bucket map[int]bool
// 	bucket := map[int]bool{}
// 	var res []int
// 	for _, v := range nums {
// 		if _, ok := bucket[v]; !ok {
// 			bucket[v] = true
// 		} else {
// 			res = append(res, v)
// 		}
// 	}
// 	return res
// }

func findDuplicates(nums []int) []int {
	var res []int
	l := len(nums)
	for _, v := range nums {
		x := (v - 1) % l
		nums[x] += l
	}
	for i, v := range nums {
		if v > (l << 1) {
			res = append(res, i+1)
		}
	}
	return res
}
