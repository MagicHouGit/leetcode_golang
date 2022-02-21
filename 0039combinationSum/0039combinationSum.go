// 39. 组合总和
// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。

// 示例 1：

// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 解释：
// 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
// 7 也是一个候选， 7 = 7 。
// 仅有这两种组合。
// 示例 2：

// 输入: candidates = [2,3,5], target = 8
// 输出: [[2,2,2,2],[2,3,3],[3,5]]
// 示例 3：

// 输入: candidates = [2], target = 1
// 输出: []

// 提示：

// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// candidate 中的每个元素都 互不相同
// 1 <= target <= 500
package main

import (
	"fmt"
)

func main() {
	c0 := []int{2, 3, 6, 7}
	t0 := 7
	c1 := []int{2, 3, 5}
	t1 := 8
	// c2 := []int{2}
	// t2 := 1
	fmt.Println(combinationSum(c0, t0))
	fmt.Println(combinationSum(c1, t1))
	// fmt.Println(combinationSum(c2, t2))
}

//===============================
//回溯加剪枝 blackTrack prune
// 执行用时：4 ms, 在所有 Go 提交中击败了56.16%的用户
// 内存消耗：4.3 MB, 在所有 Go 提交中击败了13.79的用户
//2022-02-21
//取消掉 排序之后的 速度
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：5.6 MB, 在所有 Go 提交中击败了11.60%的户

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}

	// sort.Ints(candidates)
	// invert(candidates)

	com(candidates, target, 0, []int{}, 0)
	return res

}
func com(nums []int, target int, t int, subset []int, index int) {
	if target == t {
		temp := make([]int, len(subset))
		copy(temp, subset)
		res = append(res, temp)
		return
	}
	if target < t {
		return
	}
	for i := index; i < len(nums); i++ {
		com(nums, target, t+nums[i], append(subset, nums[i]), i)
	}
}
func invert(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}

//=======================================================
//other
//这个也是回溯+剪枝,和我的方法一模一样,剪枝的方法也一样
//我比他多一步排序
// func combinationSum(candidates []int, target int) [][]int {
//     var trcak []int
//     var res [][]int
//     backtracking(0,0,target,candidates,trcak,&res)
//     return res
// }
// func backtracking(startIndex,sum,target int,candidates,trcak []int,res *[][]int){
//     //终止条件
//     if sum==target{
//         tmp:=make([]int,len(trcak))
//         copy(tmp,trcak)//拷贝
//         *res=append(*res,tmp)//放入结果集
//         return
//     }
//     if sum>target{return}
//     //回溯
//     for i:=startIndex;i<len(candidates);i++{
//         //更新路径集合和sum
//         trcak=append(trcak,candidates[i])
//         sum+=candidates[i]
//         //递归
//         backtracking(i,sum,target,candidates,trcak,res)
//         //回溯
//         trcak=trcak[:len(trcak)-1]
//         sum-=candidates[i]
//     }
// }

//================================================
//other
//over100%
//我比他多一步排序,其他一样

// func combinationSum(candidates []int, target int) (ans [][]int) {
//     // sort.Ints(candidates)
//     var backtrack func([]int, int, int)
//     backtrack = func(path []int, start int, sum int) {
//         if sum == target {
//             tmp := make([]int, len(path))
//             copy(tmp, path)
//             ans = append(ans, tmp)
//             return
//         }
//         if sum > target {
//             return
//         }
//         for i := start ; i < len(candidates); i++ {
//             path = append(path, candidates[i])
//             sum += candidates[i]
//             backtrack(path, i, sum)
//             path = path[:len(path)-1]
//             sum -= candidates[i]
//         }
//     }
//     backtrack([]int{}, 0, 0)
//     return ans
// }
