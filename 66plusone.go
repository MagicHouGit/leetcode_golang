// 66. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

// 你可以假设除了整数 0 之外，这个整数不会以零开头。

// 示例 1:

// 输入: [1,2,3]
// 输出: [1,2,4]
// 解释: 输入数组表示数字 123。
// 示例 2:

// 输入: [4,3,2,1]
// 输出: [4,3,2,2]
// 解释: 输入数组表示数字 4321。
package main

import "fmt"

func main() {
	// t_nums := []int{8, 9, 9, 9}
	// t_nums := []int{3, 6, 6, 4, 7, 5, 6, 2, 3, 8, 6, 4, 3, 2, 5, 6, 3, 5, 7, 8, 9, 5, 6, 2, 5, 6}
	t_nums := []int{9}
	// plusone(t_nums)
	// var tl int = 999/100%10
	// fmt.Println(tl)
	for _, v := range plusone(t_nums) {
		fmt.Printf("%d\n", v)
	}
}
func plusone(digits []int) []int {
	//这个方法是在我int位数不够的情况下的改进版，
	//在leecode上
	//执行用时 : 4 ms, 在Plus One的Go提交中击败了54.68% 的用户
	//内存消耗 : 2.2 MB, 在Plus One的Go提交中击败了30.57% 的用户
	size := len(digits)
	AllNine := true

	for i := 0; i < size; i++ {
		if digits[i] != 9 {
			AllNine = false
			break
		}
	}
	if AllNine == true {
		digits[0] = 1
		for i := 1; i < size; i++ {
			digits[i] = 0
		}
		digits = append(digits, 0)
		return digits
	}
	for j := size - 1; j >= 0; j-- {
		if digits[j] == 9 {
			digits[j] = 0
		} else {
			digits[j] = digits[j] + 1
			break
		}

	}

	return digits
	//这个办法被废弃了，leecode给了一个有20位的数组，我的uint64的话也只有1.8*10^19。根本不够
	// size := len(digits)
	// // var tmp int = 0
	// // for i:= 0;i<size;i++{
	// // 	// tmp += digits[i]*(10^(size-i))
	// //     tmp += digits[i]*math.Pow10(size-i)
	// //     fmt.Println(tmp)
	// // }
	// var tmp int
	// for i := 0; i < size; i++ {
	// 	tmp += digits[i] * int(math.Pow10(size-1-i))
	// }
	// // tmp := 8*math.Pow10(size)
	// Ntmp := tmp + 1
	// tmp = Ntmp
	// // fmt.Printf("%p\n", &Ntmp)
	// // fmt.Printf("%p\n", &tmp)
	// // fmt.Println(tmp)
	// var lang int
	// for tmp >= 10 {
	// 	tmp /= 10
	// 	lang++
	// }
	// // fmt.Println(lang)
	// if size-lang == 1 {
	// 	// digits[lang] = digits[lang] + 1
	// 	for i := 0; i <= lang; i++ {
	// 		digits[i] = Ntmp / int(math.Pow10(lang-i))
	// 		Ntmp = Ntmp - digits[i]*int(math.Pow10(lang-i))
	// 	}
	// }
	// if size == lang {
	// 	for i := 0; i < lang; i++ {
	// 		digits[i] = Ntmp / int(math.Pow10(lang-i))
	// 		Ntmp = Ntmp - digits[i]*int(math.Pow10(lang-i))
	// 	}
	// 	digits = append(digits, Ntmp)
	// }
	// // digits = append(digits, 3)
	// for _, v := range digits {
	// 	fmt.Printf("%d\n", v)
	// }
	// return digits

	// //当然我还是要从leecode拷贝一个其他的方法
	// var carry int = 1
	// var l = len(digits)
	// for i := l-1; i>=0; i-- {
	// 	digits[i] = digits[i] + carry
	// 	if digits[i] == 10 {
	// 		digits[i] = 0
	// 		carry =1
	// 	} else {
	// 		carry = 0
	// 	}
	// }
	// if carry == 1 {
	// 	digits = append([]int{1}, digits...)
	// }
	// return digits
}
