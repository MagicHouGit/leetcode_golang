// 40. 组合总和 II
// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]
// 提示:
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
package main

import (
	"fmt"
	"sort"
)

func main() {
	c0 := []int{10, 1, 2, 7, 6, 1, 5}
	t0 := 8
	c1 := []int{}
	t1 := 8
	c2 := []int{2, 5, 2, 1, 2}
	t2 := 5

	fmt.Println(combinationSum2(c0, t0))
	fmt.Println(combinationSum2(c1, t1))
	fmt.Println(combinationSum2(c2, t2))

}

//backTrakc
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.9 MB, 在所有 Go 提交中击败了8.98%的用户
//2022-02-27
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var backTrack func(nums []int, index int, path []int, sum int)
	backTrack = func(nums []int, index int, path []int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := index; i < len(nums); i++ {
			if nums[i] > target {
				return
			}
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			backTrack(nums, i+1, append(path, nums[i]), sum+nums[i])

		}
	}
	sort.Ints(candidates)
	backTrack(candidates, 0, []int{}, 0)
	return res
}

//backTrack
// 执行用时：4 ms, 在所有 Go 提交中击败了45.51%的用户
// 内存消耗：4.1 MB, 在所有 Go 提交中击败了8.00%的用户
// 2022-02-25
// func combinationSum2(candidates []int, target int) [][]int {
// 	var res [][]int
// 	sort.Ints(candidates)
// 	history := map[int]bool{}
// 	var backTrack func(nums []int, sum int, path []int, index int, history map[int]bool)
// 	backTrack = func(nums []int, sum int, path []int, index int, history map[int]bool) {
// 		if sum > target {
// 			return
// 		}
// 		if sum == target {
// 			temp := make([]int, len(path))
// 			copy(temp, path)
// 			res = append(res, temp)
// 			return
// 		}
// 		for i := index; i < len(nums); i++ {
// 			if nums[i] > target {
// 				return
// 			}
// 			if i > 0 && nums[i] == nums[i-1] && !history[i-1] {
// 				continue
// 			}
// 			history[i] = true
// 			backTrack(nums, sum+nums[i], append(path, nums[i]), i+1, history)
// 			history[i] = false
// 		}
// 	}
// 	backTrack(candidates, 0, []int{}, 0, history)
// 	return res
// }

//other 代码随想录
// func combinationSum2(candidates []int, target int) [][]int {
//     var trcak []int
//     var res [][]int
//     var history map[int]bool
//     history=make(map[int]bool)
//     sort.Ints(candidates)
//     backtracking(0,0,target,candidates,trcak,&res,history)
//     return res
// }
// func backtracking(startIndex,sum,target int,candidates,trcak []int,res *[][]int,history map[int]bool){
//     //终止条件
//     if sum==target{
//         tmp:=make([]int,len(trcak))
//         copy(tmp,trcak)//拷贝
//         *res=append(*res,tmp)//放入结果集
//         return
//     }
//     if sum>target{return}
//     //回溯
//     // used[i - 1] == true，说明同一树支candidates[i - 1]使用过
//     // used[i - 1] == false，说明同一树层candidates[i - 1]使用过
//     for i:=startIndex;i<len(candidates);i++{
//         if i>0&&candidates[i]==candidates[i-1]&&history[i-1]==false{
//                 continue
//         }
//         //更新路径集合和sum
//         trcak=append(trcak,candidates[i])
//         sum+=candidates[i]
//         history[i]=true
//         //递归
//         backtracking(i+1,sum,target,candidates,trcak,res,history)
//         //回溯
//         trcak=trcak[:len(trcak)-1]
//         sum-=candidates[i]
//         history[i]=false
//     }
// }
