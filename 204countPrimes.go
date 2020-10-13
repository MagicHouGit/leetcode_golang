// 统计所有小于非负整数 n 的质数的数量。

//

// 示例 1：

// 输入：n = 10
// 输出：4
// 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
// 示例 2：

// 输入：n = 0
// 输出：0
// 示例 3：

// 输入：n = 1
// 输出：0

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/count-primes
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	// test := 3
	test := 10
	fmt.Println(countPrimes(test))
}

// func countPrimes(n int) int {
// 	if n <= 2 {
// 		return 0
// 	}

// 	bo := make([]bool, n)
// 	count := 0
// 	for i := 2; i < n; i++ {
// 		if bo[i] {
// 			continue
// 		}
// 		for j := i + i; j < n; j += i {
// 			bo[j] = true
// 		}
// 		count++
// 	}
// 	return count

// }

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	count := n - 2
	isPrimes := make([]bool, n)

	for i := 2; i*i < n; i++ {
		if !isPrimes[i] {
			for j := i * i; j < n; j += i {
				if !isPrimes[j] {
					//如果当前的数没有被标记过，就将该数标记，同时count的值减1
					isPrimes[j] = true
					count--
				}
			}
		}

	}
	return count
}
