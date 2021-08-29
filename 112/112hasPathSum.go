// 112. 路径总和
// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
// 叶子节点 是指没有子节点的节点。
// 示例 1：
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// 输出：true
// 示例 2：
// 输入：root = [1,2,3], targetSum = 5
// 输出：false
// 示例 3：
// 输入：root = [1,2], targetSum = 0
// 输出：false
// 提示：
// 树中节点的数目在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
package main

import "fmt"

func main() {
	// test := []int{1, 2, 3, 0, 0, 4, 5}
	// test := []int{1, 2, 3, 4, 5, 0, 0}
	// test := []int{1}
	// test := []int{-2, 0, -3}
	test := []int{1, -2, -3, 1, 3, -1, 0, -1}
	tree := makeBinaryTree(test, 0)
	// fmt.Println(hasPathSum(tree, 8))
	// fmt.Println(hasPathSum(tree, 7))
	// fmt.Println(hasPathSum(tree, 4))
	fmt.Println(hasPathSum(tree, -4))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：4.6 MB, 在所有 Go 提交中击败了99.91%的用户
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val
	if hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) {
		return true
	}
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
func recursion(root *TreeNode, sum int) int {
	if root.Left != nil {
		recursion(root.Left, root.Val+sum)
	}
	if root.Right != nil {
		recursion(root.Right, root.Val+sum)
	}
	return sum
}

// TreeNode bbb bbb
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	newTreeNode := TreeNode{Val: x[n]}

	if 2*n+1 < l {
		newTreeNode.Left = makeBinaryTree(x, 2*n+1)
	}
	if 2*n+2 < l {
		newTreeNode.Right = makeBinaryTree(x, 2*n+2)
	}
	return &newTreeNode
}
