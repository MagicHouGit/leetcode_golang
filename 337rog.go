// 337. 打家劫舍 III
// 在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

// 示例 1:

// 输入: [3,2,3,null,3,null,1]

//      3
//     / \
//    2   3
//     \   \
//      3   1

// 输出: 7
// 解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
// 示例 2:

// 输入: [3,4,5,1,3,null,1]

//      3
//     / \
//    4   5
//   / \   \
//  1   3   1

// 输出: 9
// 解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
package main

import "fmt"

func main() {
	ia := []int{3, 2, 3, 0, 3, 0, 1}
	testTree := makeBinaryTree(ia, 0)
	fmt.Println(rob(testTree))

}

//TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	re := dfs(root)
	return max(re[0], re[1])
}
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l, r := dfs(root.Left), dfs(root.Right)
	// f 偷当前节点，g 不偷当前节点
	f := root.Val + l[1] + r[1]
	g := max(l[0], l[1]) + max(r[0], r[1])
	return []int{f, g}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func makeBinaryTree(x []int, n int) *TreeNode {

	//n初始化为0，用于递归中确定位置
	//计算层数 ，比如第一层，是0.第二层 是 1，2。第三层是 3，4，5，6.
	/*
		通过数组长度先确定层数，
		利用层数来循环
		需要记录上一层的节点，方便下一层使用（有点占内存，有没有更好的办法？）

	*/
	if x == nil || x[n] == 0 {
		return nil
	}
	l := len(x)
	// var newTreeNode *TreeNode
	// newTreeNode.Val = x[n]
	// newTreeNode.Val = 1
	newTreeNode := TreeNode{Val: x[n]}

	if 2*n+1 < l {
		newTreeNode.Left = makeBinaryTree(x, 2*n+1)
	}
	if 2*n+2 < l {
		newTreeNode.Right = makeBinaryTree(x, 2*n+2)
	}
	return &newTreeNode
}
