/*
执行用时 :4 ms, 在所有 Go 提交中击败了79.32%的用户
内存消耗 :3.5 MB, 在所有 Go 提交中击败了7.49%的用户


//*/
package main

import "fmt"

func main() {
	tmp_str := "-0012"
	fmt.Printf("%d\n", myAtoi(tmp_str))
}
func myAtoi(str string) int {

	var pn byte
	var num_s []int
	var num int
	var itr int = 0 //分为三个阶段，空白，数字，其他字符

	for _, v := range str {
		//第一阶段只能是空白

		if itr == 0 {
			//空白直接下一轮 空格ASCII位置32，
			if v == 32 {
				continue
			}

			//第一个正负号或者数字
			if v >= '0' && v <= '9' {
				num_s = append(num_s, int(v)-'0')
				itr = 1
				continue
			}
			if v == '-' || v == '+' {
				pn = byte(v)
				itr = 1
				continue
			} else {
				//先出现了数字正负号的其他字符，不然换，return 0.
				return 0
			}
		}
		//正负号已经存入，或没有正负号。进入数字阶段
		if itr == 1 {
			if v >= '0' && v <= '9' {
				num_s = append(num_s, int(v)-'0')

			} else {
				itr = 2
				break
			}
		}
	}
	if itr != 2 {
		fmt.Errorf("数据没有完整遍历")
	}
	for i := 0; i < len(num_s); i++ {
		if i == 0 {
			num = int(num_s[i])
			if pn == '-' {
				num = num * -1
			}
		}
		if pn == '-' {
			if num*10 <= -2147483648 {
				return -2147483648
			}
			if num == -214748364 && int(num_s[i]) > 8 {
				return -2147483648
			}
		} else {
			if num*10 >= 2147483647 {
				return 2147483647
			}
			if num == 214748364 && int(num_s[i]) > 7 {
				return 2147483647
			}

		}
		if i > 0 {
			if pn == '-' {
				num = num*10 - int(num_s[i])
			} else {
				num = num*10 + int(num_s[i])
			}
		}

	}

	return num
}
