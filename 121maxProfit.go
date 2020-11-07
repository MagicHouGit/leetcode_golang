// 121. 买卖股票的最佳时机
// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

// 注意：你不能在买入股票前卖出股票。

// 示例 1:

// 输入: [7,1,5,3,6,4]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 示例 2:

// 输入: [7,6,4,3,1]
// 输出: 0
// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
package main

import "fmt"

func main() {
	// test := []int{7, 1, 5, 3, 6, 4}
	test := []int{1, 5}
	fmt.Println(maxProfit(test))

}

// 执行用时：216 ms, 在所有 Go 提交中击败了14.63%的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了20.42%的用户
func maxProfit(prices []int) int {
	re := 0
	l := len(prices)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			pre := prices[j] - prices[i]
			if pre > re {
				re = pre
			}
		}
	}
	return re
}

// 这个是速度最快的一层for，O(n)
// func maxProfit(prices []int) int {
// 	var ans, mm = 0, 0
// 	for i := len(prices) - 1; i >= 0; i-- {
// 		if prices[i] > mm {
// 			mm = prices[i]
// 		}
// 		d := mm - prices[i]
// 		if d > ans {
// 			ans = d
// 		}
// 	}
// 	return ans
// }
