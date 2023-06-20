// 95. 不同的二叉搜索树 II
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
// 示例 1：
// 输入：n = 3
// 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
// 示例 2：
// 输入：n = 1
// 输出：[[1]]
// 提示：
// 1 <= n <= 8

package main

import "fmt"

func main() {
	// h3 := &TreeNode{Val: 3}
	// h2 := &TreeNode{Val: 2, Left: h3}
	// h1 := &TreeNode{Val: 1, Right: h2}
	// h0 := &TreeNode{0, nil, h1}
	// treeStd(h0)
	// treeStd(h1)

	n3 := 3
	head := generateTrees(n3)
	for _, v := range head {
		treeStd(v)
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode
	var top *TreeNode
	// var temp *TreeNode

	var backTrack func(n int, l int, nums []int, head *TreeNode)
	backTrack = func(n int, l int, nums []int, head *TreeNode) {
		if l == n {
			t := top
			res = append(res, t)
			// top = nil
			return
		}

		for i := 0; i < len(nums); i++ {
			temp := &TreeNode{nums[i], nil, nil}
			if l == 0 {
				top = temp
			}

			take := nums[i]
			nums = append(nums[:i], nums[i+1:]...)

			if head.Val < temp.Val {
				head.Right = temp
				backTrack(n, l+1, nums, head.Right)
				// head.Right = nil
			} else {
				head.Left = temp
				backTrack(n, l+1, nums, head.Left)
				// head.Right = nil
			}
			nums = append(nums[:i], append([]int{take}, nums[i:]...)...)

		}
	}
	// top := &TreeNode{0, nil, nil}
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	backTrack(n, 0, nums, &TreeNode{})
	return res

}

func treeStd(head *TreeNode) {
	var res [][]int
	res = append(res, []int{0})
	// for i := 0; i <= 3; i++ {
	// 	temp := make([]int, 1<<i)
	// 	res = append(res, temp)
	// }
	var bfs func(h *TreeNode, site int, layer int)
	bfs = func(h *TreeNode, site int, layer int) {
		if h == nil {
			return
		}
		if len(res) == layer+1 {
			temp := make([]int, 1<<(layer+1))
			res = append(res, temp)
		}
		res[layer][site] = h.Val
		if h.Left != nil {
			bfs(h.Left, site<<1, layer+1)
		}
		if h.Right != nil {
			bfs(h.Right, (site<<1)+1, layer+1)
		}
	}
	bfs(head, 0, 0)
	fmt.Println(res)
}

// func bfs(head *TreeNode, site int, layer int, res [][]int) {
// 	if *TreeNode == nil {
// 		return
// 	}
// 	res[layer][site] = head.Val

// 	bfs(h)

// }

// ==================================================
// //toher
// var dp = [10][10][]*TreeNode{}
// func generateTrees(n int) []*TreeNode {
//     return dfs(1, n)
// }
// func dfs(i, j int) []*TreeNode {
//     if dp[i][j] != nil {
//         return dp[i][j]
//     }
//     ret := []*TreeNode{}
//     if j < i {
//         return append(ret, nil)
//     }
//     if i == j {
//         dp[i][j] = append(ret, &TreeNode{Val: i})
//         return dp[i][j]
//     }
//     for k := i; k <= j; k++ {
//         l, r := dfs(i, k - 1), dfs(k + 1, j)
//         for _, lnode := range l {
//             for _, rnode := range r {
//                 ret = append(ret, &TreeNode{k, lnode, rnode})
//             }
//         }
//     }
//     dp[i][j] = ret
//     return ret
// }
