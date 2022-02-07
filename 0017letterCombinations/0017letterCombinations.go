package main

import "fmt"

func main() {
	// x :=
	// d0 := "2"
	// d1 := "23"
	d2 := "237"
	// fmt.Println(letterCombinations(d0))
	// fmt.Println(letterCombinations(d1))
	fmt.Println(letterCombinations(d2))

}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了94.38%的用户
// 速度还可以,还是要写一版回溯的解决方案
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// func letterCombinations(digits string) []string {
// 	res := []string{}
// 	for i := 0; i < len(digits); i++ {
// 		res = quanpailie(string(digits[i]), res)
// 	}
// 	return res
// }

// func quanpailie(digits string, temp []string) []string {
// 	res := []string{}
// 	if len(temp) == 0 {
// 		for i := 0; i < len(phoneMap[digits]); i++ {
// 			res = append(res, string(phoneMap[digits][i]))
// 		}
// 	}
// 	for i := 0; i < len(temp); i++ {
// 		for ii := 0; ii < len(phoneMap[digits]); ii++ {
// 			res = append(res, temp[i]+string(phoneMap[digits][ii]))
// 		}
// 	}
// 	return res
// }

//================================
// backTrack
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了94.50%的用户
var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backTrack(digits, 0, "")
	return combinations
}

func backTrack(digits string, index int, combination string) {
	if len(digits) == index {
		combinations = append(combinations, combination)
	} else {
		nu := string(digits[index])
		letters := phoneMap[nu]
		for i := 0; i < len(letters); i++ {
			backTrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

//==================================
//other

// var combinations []string

// func letterCombinations(digits string) []string {
// 	if len(digits) == 0 {
// 		return []string{}
// 	}
// 	combinations = []string{}
// 	backtrack(digits, 0, "")
// 	return combinations
// }

// func backtrack(digits string, index int, combination string) {
// 	if index == len(digits) {
// 		combinations = append(combinations, combination)
// 	} else {
// 		digit := string(digits[index])
// 		letters := phoneMap[digit]
// 		lettersCount := len(letters)
// 		for i := 0; i < lettersCount; i++ {
// 			backtrack(digits, index+1, combination+string(letters[i]))
// 		}
// 	}
// }
