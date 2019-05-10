package main

import "fmt"

func main() {
	t_nums := []int{0, 1, 0, 3, 12}
	moveZeroes(t_nums)
	for _, v := range t_nums {
		fmt.Printf("%d\n", v)
	}
	// for _,v range t_nums {
	// 	fmt.Printf("%d\n",v)
	// }

}

//换位之法，拿到一个0，往后找一个不是0的，互换，循环一边结束
func moveZeroes(nums []int) {
	lastZero := len(nums) - 1
	for i := 0; i <= lastZero; i++ {
		if nums[i] != 0 {
			continue
		} else {
			for j := i + 1; j <= lastZero; j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				} else {
					continue
				}
			}
		}

	}

}

// 此方法来自leecode，比我的快
// func moveZerose (nums []int){
// 	o := 0
// 	for i := 0; i < len(nums); i++ {
// 		v := nums[i]
// 		if v == 0 {
// 			continue
// 		} else {
// 			nums[o] = v
// 			o++
// 		}
// 	}
// 	for i := o; i < len(nums); i++ {
// 		nums[i] = 0
// 	}
// }
