package main

import "fmt"

func main() {
	// test := "III"
	// test1 := "01234"
	// test := "IV"
	// test := "V"
	// test := "LVIII"
	// test := "LIV"
	test := "MCMXCIV"
	// fmt.Println(test1[0:2])
	// fmt.Println(test[2])
	// fmt.Printf("%s\n", test[1])
	// fmt.Printf("%c\n", test[1])
	// fmt.Printf("%p\n", test[1])
	fmt.Println(romanToInt(test))
}

// 执行用时：16 ms, 在所有 Go 提交中击败了23.32%的用户
// 内存消耗：5 MB, 在所有 Go 提交中击败了5.03%的用户
// func romanToInt(s string) int {
// 	if s == "" {
// 		return 0
// 	}
// 	index := make(map[string]int)
// 	index["I"] = 1
// 	index["V"] = 5
// 	index["X"] = 10
// 	index["L"] = 50
// 	index["C"] = 100
// 	index["D"] = 500
// 	index["M"] = 1000
// 	index["IV"] = 4
// 	index["IX"] = 9
// 	index["XL"] = 40
// 	index["XC"] = 90
// 	index["CD"] = 400
// 	index["CM"] = 900
// 	var res int
// 	for i := 0; i < len(s); {
// 		if len(s) < 2 {
// 			break
// 		}
// 		_, ok := index[s[0:2]]
// 		if ok {
// 			res += index[string(s[0:2])]
// 			s = s[2:]
// 		} else {
// 			res += index[string(s[0])]
// 			s = s[1:]
// 		}
// 	}
// 	if len(s) == 1 {
// 		res += index[string(s[0])]
// 	}

// 	return res
// }

//----------------------
func romanToInt(s string) (r int) {

	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := judge(s[i])
		if cur == -1 {
			return -1
		} else {
			if cur >= pre {
				r += cur
			} else {
				r -= cur
			}
			pre = cur
		}

	}
	return r
}

func judge(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return -1
	}
}
