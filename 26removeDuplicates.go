//从排序数组中删除重复项
package main

import (
	"fmt"
)

func main() {
	numsArr := []int{1, 1, 2, 2, 2, 3}
	// fmt.Printf("%d\n",len(nums))
	removeDuplicates(numsArr)
}
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	max := nums[len(nums)-1]
	for i := 1; i <= max; i++ {
		nums[i] = nums[0] + i
	}
	for _, value := range nums {

		fmt.Printf("%d", value)
	}

	return max - nums[0] + 1
}

//这段不知道为什么在leecode上的答案是有问题的，默认最后一位总是1
// func removeDuplicates(nums []int) int {
//   if len(nums) == 1 {
//     fmt.Printf("%d",nums[0])
//     return 1
//   }
//   //Dnums用来接收过滤后的数组
//   Dnums:=[]int{}
//   //true 表示不是重复的数，false表示重复的数
//   tmp := true
//   //先添加第一个数（肯定不是重复的）
//   Dnums = append(Dnums,nums[0])
//   for oi:=1;oi<len(nums);oi++{
//     tmp = true
//    //从nums选一个数在Snums中查看是否已经存在
//     for pi:=0;pi<len(Dnums);pi++{
//       if nums[oi]== Dnums[pi]{
//         tmp =false
//         break
//       }
//     }
//     //这是一个不重复的数就添加
//     if tmp == true {
//       Dnums = append(Dnums,nums[oi])
//       //tmp = false
//     }
//   }

//   for _,value := range Dnums {
//     fmt.Printf("%d\n",value)
//   }
//   return len(Dnums)
//}
