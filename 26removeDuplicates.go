// 26. 删除排序数组中的重复项
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

// 示例 1:

// 给定数组 nums = [1,1,2],

// 函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

// 你不需要考虑数组中超出新长度后面的元素。
// 示例 2:

// 给定 nums = [0,0,1,1,1,2,2,3,3,4],

// 函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

// 你不需要考虑数组中超出新长度后面的元素。

// 说明:

// 为什么返回数值是整数，但输出的答案是数组呢?

// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

// 你可以想象内部操作如下:

// // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
// int len = removeDuplicates(nums);

// // 在函数里修改输入数组对于调用者是可见的。
// // 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
// for (int i = 0; i < len; i++) {
//     print(nums[i]);
// }
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
