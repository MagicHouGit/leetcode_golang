// 98. 验证二叉搜索树
// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 假设一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 示例 1:

// 输入:
//     2
//    / \
//   1   3
// 输出: true
// 示例 2:

// 输入:
//     5
//    / \
//   1   4
//      / \
//     3   6
// 输出: false
// 解释: 输入为: [5,1,4,null,null,3,6]。
//      根节点的值为 5 ，但是其右子节点值为 4 。
package main

import (
	"fmt"
	"math"
)

func main() {
	// testArrTree := []int{1, 2, 3, 0, 5, 0, 7, 0, 0, 10, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 5, 0, 7, 0, 0, 10, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 0, 6, 0}
	// testArrTree := []int{2, 1, 3}
	// testArrTree := []int{1, 0, 3, 0, 0, 4, 0}
	// testArrTree := []int{3, 1, 4, 0, 0, 5, 6}
	// testArrTree := []int{3, 1, 5, 0, 0, 4, 6}

	testArrTree := []int{1, 2, 3, 4, 5, 0, 6, 0, 0, 7, 8}
	// testArrTree := []int{5, 1, 4, 0, 0, 3, 6}
	// testArrTree := []int{0}
	// testArrTree := []int{5, 1, 4, 0, 0, 3, 6}
	testTree := makeBinaryTree(testArrTree, 0)
	// testTree := &TreeNode{Val: 0}
	fmt.Println(middleTraversal(testTree))
	// fmt.Println(isValidBST(testTree))
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)
}

// TreeNode struct {
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// var gt []int
// 执行用时：36 ms, 在所有 Go 提交中击败了14.16%的用户
// 内存消耗：5.9 MB, 在所有 Go 提交中击败了10.46%的用户
func middleTraversal(root *TreeNode) []int {

	var res []int
	var stack []*TreeNode
	// var temp *TreeNode
	if root == nil {
		return res
	}
	for root != nil || stack != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if stack != nil {
			Le := len(stack)
			if Le == 0 {
				break
			}
			root = stack[Le-1]
			stack = stack[:Le-1]
			res = append(res, root.Val)
			root = root.Right
		}

	}
	return res

}

// func isValidBST(root *TreeNode) bool {

// 	// middleTraversal(root)
// 	gt := middleTraversal(root)
// 	for i := 0; i < len(gt)-1; i++ {
// 		for j := i + 1; j < len(gt); j++ {

// 			if gt[i] >= gt[j] {
// 				return false
// 			}
// 		}
// 	}
// 	return true

// }
///这个是最快的
func isValidBST(root *TreeNode) bool {
	return help(root, math.MinInt64, math.MaxInt64)
}

func help(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return help(root.Left, min, root.Val) && help(root.Right, root.Val, max)
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
