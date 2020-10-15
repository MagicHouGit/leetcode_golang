package main

import "fmt"

func main() {
	// test := 200
	// test := 3
	// test := 4
	// test := 9
	// test := 58
	test := 1994
	fmt.Println(intToRoman(test))
}

// 执行用时：8 ms, 在所有 Go 提交中击败了84.04%的用户
// 内存消耗：3.4 MB, 在所有 Go 提交中击败了45.20%的用户

// func intToRoman(num int) string {

// 	arr := [4]int{}

// 	arr[0] = num / 1000
// 	arr[1] = (num - arr[0]*1000) / 100
// 	arr[2] = (num - arr[0]*1000 - arr[1]*100) / 10
// 	arr[3] = num - arr[0]*1000 - arr[1]*100 - arr[2]*10
// 	var a, b, c string
// 	var res string
// 	for i := 0; i < 4; i++ {
// 		a, b, c = re(i)
// 		// re = append(re,nm(a,b,c,arr[i]))
// 		res = res + nm(a, b, c, arr[i])
// 	}
// 	return res

// }
// func nm(a string, b string, c string, n int) string {
// 	switch n {
// 	case 1:
// 		return a
// 	case 2:
// 		return a + a
// 	case 3:
// 		return a + a + a
// 	case 4:
// 		return a + b
// 	case 5:
// 		return b
// 	case 6:
// 		return b + a
// 	case 7:
// 		return b + a + a
// 	case 8:
// 		return b + a + a + a
// 	case 9:
// 		return a + c
// 	}
// 	return ""
// }

// func re(n int) (string, string, string) {
// 	switch n {
// 	case 0:
// 		return "M", "", ""
// 	case 1:
// 		return "C", "D", "M"
// 	case 2:
// 		return "X", "L", "C"
// 	case 3:
// 		return "I", "V", "X"
// 	}
// 	return "", "", ""
// }
func intToRoman(num int) string {
	var arr []int
	var div, rem int
	var ret string
	const_5 := []string{"V", "L", "D"}
	const_1 := []string{"I", "X", "C", "M"}

	getS := func(num, digit int) string {
		if num == 0 {
			return ""
		}
		if num == 9 {
			return const_1[digit] + const_1[digit+1]
		}
		if num == 4 {
			return const_1[digit] + const_5[digit]
		}
		if num == 5 {
			return const_5[digit]
		}
		if num <= 3 {
			tmp := ""
			for i := 1; i <= num; i++ {
				tmp += const_1[digit]
			}
			return tmp
		}

		tmp := const_5[digit]
		for i := 6; i <= num; i++ {
			tmp += const_1[digit]
		}
		return tmp
	}

	for {
		if num == 0 {
			break
		}
		div, rem = num/10, num%10
		arr = append(arr, rem)
		num = div
	}

	for i := len(arr) - 1; i >= 0; i-- {
		ret = ret + getS(arr[i], i)
	}
	return ret
}
