// 306. 累加数
// 累加数 是一个字符串，组成它的数字可以形成累加序列。
// 一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，序列中的每个后续数字必须是它之前两个数字之和。
// 给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。
// 说明：累加序列里的数，除数字 0 之外，不会 以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

// 示例 1：
// 输入："112358"
// 输出：true
// 解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8

// 示例 2：
// 输入："199100199"
// 输出：true
// 解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199

// 提示：
// 1 <= num.length <= 35
// num 仅由数字（0 - 9）组成

// 进阶：你计划如何处理由过大的整数输入导致的溢出?
package main

import "fmt"

func main() {

	str0 := "1910"
	str1 := "199100199"
	str2 := "12122436"
	str3 := "199001200"
	str4 := "011235"
	fmt.Println(isAdditiveNumber(str0)) //true
	fmt.Println(isAdditiveNumber(str1)) //true
	fmt.Println(isAdditiveNumber(str2)) //true
	fmt.Println(isAdditiveNumber(str3)) //false
	fmt.Println(isAdditiveNumber(str4)) //true
	// t0 := "8"
	// t1 := "16"
	// fmt.Println(stringAdd(t0, t1))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了99.67%的用户
//2022-02-17
func isAdditiveNumber(num string) bool {

	// var backTrack func(firstEnd,secondEnd int,num string,jg bool)

	return backTrack(num, 1)

}

//单个数字位数不能超过总长度的三分之一,
//first长度不能大于二分之一,second长度不能超过剩余的二分之一
//third长度大于等于first和second最长的长度
func backTrack(num string, startLen int) bool {
	l := len(num)
	for i := startLen; i <= l/2; i++ {
		yu := len(num) - i
		for ii := 1; ii <= yu/2; ii++ {
			third := stringAdd(num[:i], num[i:ii+i])
			if len(num) < len(third)+i+ii {
				break
			}
			if (i > 2 && num[0] == '0') || (ii > 2 && num[i] == '0') {
				break
			}
			if third == num[ii+i:ii+i+len(third)] {
				if l == i+ii+len(third) {
					return true
				}
				res := backTrack(num[i:], ii)
				if res == true {
					return true
				}
			}
		}
	}
	return false
}
func stringAdd(x, y string) string {
	byteX := []byte(x)
	byteY := []byte(y)
	longer := len(y)
	if len(x) > len(y) {
		longer = len(x)
	}
	temp := 0
	res := []byte{}
	for i := 0; i < longer; i++ {
		if len(byteX) != 0 {
			temp = temp + int(byteX[len(byteX)-1]-'0')
			byteX = byteX[:len(byteX)-1]
		}
		if len(byteY) != 0 {
			temp = temp + int(byteY[len(byteY)-1]-'0')
			byteY = byteY[:len(byteY)-1]
		}
		res = append(res, byte(temp%10+'0'))
		temp = temp / 10
	}
	if temp > 0 {
		res = append(res, byte(temp+'0'))
	}
	res = inversed(res)
	return string(res)

}
func inversed(s []byte) []byte {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return s
}

//==========================================
//other
// func stringAdd(x, y string) string {
// 	res := []byte{}
// 	carry, cur := 0, 0
// 	for x != "" || y != "" || carry != 0 {
// 		cur = carry
// 		if x != "" {
// 			cur += int(x[len(x)-1] - '0')
// 			x = x[:len(x)-1]
// 		}
// 		if y != "" {
// 			cur += int(y[len(y)-1] - '0')
// 			y = y[:len(y)-1]
// 		}
// 		carry = cur / 10
// 		cur %= 10
// 		res = append(res, byte(cur)+'0')
// 	}
// 	for i, n := 0, len(res); i < n/2; i++ {
// 		res[i], res[n-1-i] = res[n-1-i], res[i]
// 	}
// 	return string(res)
// }

// func valid(num string, secondStart, secondEnd int) bool {
// 	n := len(num)
// 	firstStart, firstEnd := 0, secondStart-1
// 	for secondEnd <= n-1 {
// 		third := stringAdd(num[firstStart:firstEnd+1], num[secondStart:secondEnd+1])
// 		thirdStart := secondEnd + 1
// 		thirdEnd := secondEnd + len(third)
// 		if thirdEnd >= n || num[thirdStart:thirdEnd+1] != third {
// 			break
// 		}
// 		if thirdEnd == n-1 {
// 			return true
// 		}
// 		firstStart, firstEnd = secondStart, secondEnd
// 		secondStart, secondEnd = thirdStart, thirdEnd
// 	}
// 	return false
// }

// func isAdditiveNumber(num string) bool {
// 	n := len(num)
// 	for secondStart := 1; secondStart < n-1; secondStart++ {
// 		if num[0] == '0' && secondStart != 1 {
// 			break
// 		}
// 		for secondEnd := secondStart; secondEnd < n-1; secondEnd++ {
// 			if num[secondStart] == '0' && secondStart != secondEnd {
// 				break
// 			}
// 			if valid(num, secondStart, secondEnd) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

//============================
// func isAdditiveNumber(num string) bool {
// 	var isCanAdded func(num string, first, second, sumIdx int) bool
// 	isCanAdded = func(num string, first, second, sumIdx int) bool {
// 		if sumIdx == len(num) {
// 			return true
// 		}
// 		sumStr := strconv.Itoa(first + second)
// 		if sumIdx+len(sumStr) > len(num) {
// 			return false
// 		}

// 		actualSum := num[sumIdx : sumIdx+len(sumStr)]
// 		return sumStr == actualSum && isCanAdded(num, second, first+second, sumIdx+len(sumStr))
// 	}

// 	first := 0
// 	for i := 0; i < len(num); i++ {
// 		if i > 0 && num[0] == '0' {
// 			return false
// 		}

// 		first = first*10 + int(num[i]-'0')
// 		second := 0
// 		for j := i + 1; j < len(num); j++ {
// 			second = second*10 + int(num[j]-'0')
// 			if j > i+1 && num[i + 1] == '0' {
// 				break
// 			}
// 			if j+1 < len(num) && isCanAdded(num, first, second, j+1) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// 作者：kyushu
// 链接：https://leetcode-cn.com/problems/additive-number/solution/rustgolangjava-hui-su-yu-dfs-100jie-fa-b-y8xm/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
