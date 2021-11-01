package singleNumber

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums1 := []int{1, 2, 1, 3, 2, 5}
	nums2 := []int{0, -1}
	nums3 := []int{1, 1, 0, -2147483648}
	fmt.Println(singleNumber(nums1))
	fmt.Println(singleNumber(nums2))
	fmt.Println(singleNumber(nums3))
	fmt.Println(singleNumber1(nums1))
	fmt.Println(singleNumber1(nums2))
	fmt.Println(singleNumber1(nums3))
}
