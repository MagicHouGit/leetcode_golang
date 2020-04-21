package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 3}
	fmt.Println(arrayPairSum(arr))

}
func arrayPairSum(nums []int) int {
	res := 0
	nums = merge(nums)
	for i := 0; i < len(nums); i++ {
		if i%2 == 1 {
			res += nums[i]
		}
	}
	return res
}
func merge(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	half := length / 2
	left := merge(arr[:half])
	right := merge(arr[half:])
	return mergeSort(left, right)
}
func mergeSort(left []int, right []int) (result []int) {
	l, r := 0, 0
	for len(left) < l && len(right) < r {
		if left[l] > right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[:l]...)
	result = append(result, right[:r]...)
	return
}

//这个是速度最快的那个题解
//暂时没有看懂
func arrayPairSum(nums []int) int {
	var buckets [20001]int8
	for _, num := range nums {
		buckets[num+10000]++
	}

	sum := 0
	needAdd := true

	for num, count := range buckets {
		for count > 0 {
			if needAdd {
				sum += num - 10000
			}
			needAdd = !needAdd
			count--
		}
	}

	return sum
}
