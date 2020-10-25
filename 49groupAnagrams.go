// 49. 字母异位词分组
// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

// 示例:

// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// 说明：

// 所有输入均为小写字母。
// 不考虑答案输出的顺序。

package main

import (
	"fmt"
)

func main() {
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// test := []string{"", "t"}
	// test := []string{"", ""}

	// a := "eat"
	// b := "tan"
	// a := "ttit"
	// b := "ittt"
	// test := []string{"hhhhu", "tttti", "tttit", "hhhuh", "hhuhh", "tittt"}
	// test := []string{"ttttit", "ittttt"}
	// fmt.Println(groupAnagrams(test))
	// fmt.Println(isAnagram(a, b))
	// fmt.Println(test)
	res := groupAnagrams(test)
	for _, v := range res {
		fmt.Println(v)
	}

}

// type by []byte

// // 自定义排序规则
// func (b by) Len() int {
// 	return len(b)
// }
// func (b by) Less(i, j int) bool {
// 	return b[i] < b[j]
// }
// func (b by) Swap(i, j int) {
// 	b[i], b[j] = b[j], b[i]
// }
// func groupAnagrams(strs []string) [][]string {
// 	ret := [][]string{}
// 	m := make(map[string]int)
// 	for _, str := range strs {
// 		uByte := by(str)
// 		sort.Sort(uByte) // 将字符串排序
// 		// sort.Sort([]byte{str})
// 		k := string(uByte)
// 		if idx, ok := m[k]; !ok {
// 			m[k] = len(ret) // 记录拍完序的字符串以及字符串在ret的位置
// 			ret = append(ret, []string{str})
// 		} else { // 已经出现过，将其放入出现在ret中，在ret的位置为idx
// 			ret[idx] = append(ret[idx], str)
// 		}
// 	}
// 	return ret
// }

//============================================
// 执行用时：520 ms, 在所有 Go 提交中击败了5.31%的用户
// 内存消耗：8.3 MB, 在所有 Go 提交中击败了28.28%的用户
// func groupAnagrams(strs []string) [][]string {

// 	var t bool
// 	var res [][]string
// 	temp := []string{}
// 	l := len(strs)
// 	if l < 1 {
// 		return nil
// 	}
// 	temp = append(temp, strs[0])
// 	res = append(res, temp)
// 	for i := 1; i < l; i++ {
// 		t = false
// 		for ii := 0; ii < len(temp); ii++ {
// 			if isAnagram(temp[ii], strs[i]) {
// 				res[ii] = append(res[ii], strs[i])
// 				t = true
// 				break
// 			}
// 		}
// 		if t == false {

// 			temp = append(temp, strs[i])
// 			res = append(res, []string{strs[i]})
// 		}

// 	}
// 	return res
// }
// func isAnagram(a, b string) bool {

// 	l0 := len(a)
// 	l1 := len(b)
// 	if l0 != l1 {
// 		return false
// 	}
// 	isExist := false
// 	for i := 0; i < l1; i++ {
// 		isExist = false
// 		for ii := 0; ii < len(a); ii++ {
// 			if b[i] == a[ii] {
// 				isExist = true
// 				a = a[0:ii] + a[ii+1:]
// 				break
// 			}
// 		}
// 		if isExist == false {
// 			return false
// 		}
// 	}
// 	return true
// }
var list = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func groupAnagrams(strs []string) [][]string {
	product2Index := map[int]int{}
	index := 0
	rets := make([][]string, 0, len(strs))
	for i := 0; i < len(strs); i++ {
		product := helper(strs[i])
		if k, exist := product2Index[product]; exist {
			rets[k] = append(rets[k], strs[i])
		} else {
			rets = append(rets, []string{strs[i]})
			product2Index[product] = index
			index++
		}
	}

	return rets
}

func helper(a string) int {
	ret := 1
	for i := 0; i < len(a); i++ {
		ret *= list[int(a[i]-'a')]
	}
	return ret
}
