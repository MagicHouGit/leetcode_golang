// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

// 说明: 叶子节点是指没有子节点的节点。

// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最大深度 3 。
package main

import "fmt"

func main() {
	// testArrTree := []int{1, 2, 3, 0, 5, 0, 7, 0, 0, 10, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 14, 15}
	testArrTree := []int{1, 2, 3, 0, 5, 0, 7, 0, 0, 10, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 0, 6, 0}
	// testArrTree := []int{1, 2, 3}
	// testArrTree := []int{1, 0, 3, 0, 0, 4, 0}
	testTree := makeBinaryTree(testArrTree, 0)
	fmt.Println(maxDepth(testTree))
}

// TreeNode struct {
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：4.5 MB, 在所有 Go 提交中击败了12.58%的用户
func maxDepth(root *TreeNode) int {

	var res int = 0
	var LT int = 0
	var RT int = 0
	if root == nil {
		return res
	} else {
		res++
	}
	if root.Left != nil {
		LT = maxDepth(root.Left)
	}
	if root.Right != nil {
		RT = maxDepth(root.Right)
	}
	if LT > RT {
		res = res + LT
	} else {
		res = res + RT
	}
	return res
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
