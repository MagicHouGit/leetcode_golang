package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// time0 := []string{"23:59", "00:00"}
	// time1 := []string{"00:00", "23:59", "00:00"}
	time2 := []string{"05:31", "22:08", "00:35"}
	// ss := strings.Split(time0[0], ":")
	// fmt.Println(ss[0])
	// fmt.Println(strings.Split(time0[0], ":"))
	// fmt.Printf("%T\n", ss)
	// fmt.Printf("%T\n", ss[0])

	// fmt.Println(findMinDifference(time0))
	// fmt.Println(findMinDifference(time1))
	fmt.Println(findMinDifference(time2))

}

// 执行用时：244 ms, 在所有 Go 提交中击败了7.46%的用户
// 内存消耗：6.2 MB, 在所有 Go 提交中击败了10.45%的用户
// 看到这时间我都惊了,
// 看了评论加了
// if len(timePoints) > 1440 {
// 	return 0
// }

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：4.3 MB, 在所有 Go 提交中击败了68.66%的用户
//这什么测试用例,我想骂人
//我还加了这个呢
// for minute > 1440 {
// 	minute -= 1440
// }

func findMinDifference(timePoints []string) int {
	var agg []int
	for i := 0; i < len(timePoints); i++ {
		temp := strings.Split(timePoints[i], ":")

		h, _ := strconv.Atoi(temp[0])
		m, _ := strconv.Atoi(temp[1])
		minute := h*60 + m

		for minute > 1440 {
			minute -= 1440
		}
		if minute == 0 {
			minute = 1440
		}
		agg = append(agg, minute)
	}
	var res = 720
	for i := 0; i < len(agg)-1; i++ {
		for ii := i + 1; ii < len(agg); ii++ {
			c := agg[i] - agg[ii]
			if c < 0 {
				c = -1 * c
			}
			if c > 720 {
				c = 1440 - c
			}
			if c < res {
				res = c
			}
		}
	}
	return res
}
