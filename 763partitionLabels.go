package main

import "fmt"

func main() {

	// S := "ababcbacadefegdehijhklij"
	S := "ab"
	// S := "abc"
	// S := "abac"
	// S := "a"
	// S := ""
	// S := "aa"
	// S := "abab"
	// S := "caedbdedda"
	// to := make(map[string][2]int, 26)
	// fmt.Println(to)
	fmt.Println(len(S))
	fmt.Println(partitionLabels(S))
}

//调试花了不少时间，a,ab,abc,abac,这个几种情况不能兼顾
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了20.00%的用户

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return nil
	}

	var repeat []byte
	var to [][]int
	var res []int

	repeat = append(repeat, S[0])
	to = append(to, []int{0, 0})
	for i, v := range S {
		for k := 0; k < len(repeat); k++ {
			if byte(v) == repeat[k] {
				to[k][1] = i
				break
			}
			if k == len(repeat)-1 {
				repeat = append(repeat, byte(v))
				to = append(to, []int{i, i})
				break
			}
		}
	}
	fmt.Println(repeat)
	fmt.Println(to)

	remain := len(S)
	start, end := to[0][0], to[0][1]
	for i := 1; i < len(repeat); i++ {

		// if to[i][1] == 0 {
		// 	// if to[i][0] == 0 {
		// 	// 	res = append(res, 1)
		// 	// 	start, end = 1, 1
		// 	// 	continue
		// 	// }
		// 	to[i][1] = to[i][0]
		// }
		// if start == end && i < len(repeat)-1 {
		// 	res = append(res, 1)
		// 	start = to[i+1][0]
		// 	end = to[i+1][1]
		// 	continue
		// }
		if to[i][0] > end {
			res = append(res, end-start+1)
			remain -= end - start + 1
			start = to[i][0]
			end = to[i][1]
			continue
		}
		if to[i][1] > end {
			end = to[i][1]
		}
		// if i == len(repeat)-1 {
		// 	res = append(res, end-start+1)
		// }
	}

	if remain > 0 {
		res = append(res, remain)
	}
	return res
}
