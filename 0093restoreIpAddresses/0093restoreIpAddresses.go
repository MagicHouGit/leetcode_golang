// 93. 复原 IP 地址
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
// 示例 1：
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 示例 2：
// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 示例 3：
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
// 提示：
// 1 <= s.length <= 20
// s 仅由数字组成

package main

import (
	"fmt"
)

func main() {

	s0 := "255255255255"
	s1 := "25525511135"
	s2 := "2550"
	s3 := "101023"
	fmt.Println(restoreIpAddresses(s0))
	fmt.Println(restoreIpAddresses(s1))
	fmt.Println(restoreIpAddresses(s2))
	fmt.Println(restoreIpAddresses(s3))

}

//backTrack
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了6.70%的用户
// 2022-03-05
//
func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {
		return []string{}
	}
	var res []string

	var backTrack func(cut []string, index int, layer int)
	backTrack = func(cut []string, index int, layer int) {
		if layer == 4 && index == len(s) {
			temp := cut[0] + "." + cut[1] + "." + cut[2] + "." + cut[3]
			res = append(res, temp)
		}

		for i := index; i < len(s); i++ {
			if !isValid(s[index : i+1]) {
				return
			}
			backTrack(append(cut, s[index:i+1]), i+1, layer+1)

		}
	}
	backTrack([]string{}, 0, 0)
	return res

}

func isValid(s string) bool {
	b := []byte(s)
	if len(b) > 1 && b[0] == '0' {
		return false
	}

	num := 0
	var temp int
	for i := 0; i < len(s); i++ {
		temp = int(b[i] - '0')
		num = num*10 + temp
	}
	return num <= 255
	// if num > 255 {
	// 	return false
	// }
	// return true
}
