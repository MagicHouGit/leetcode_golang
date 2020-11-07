// 101. 对称二叉树
// 给定一个二叉树，检查它是否是镜像对称的。

// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3

// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

//     1
//    / \
//   2   2
//    \   \
//    3    3

// 进阶：

// 你可以运用递归和迭代两种方法解决这个问题吗？

package main

import (
	"fmt"
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

	// testArrTree := []int{1, 2, 3, 4, 5, 0, 6, 0, 0, 7, 8}
	// testArrTree := []int{5, 1, 4, 0, 0, 3, 6}
	// testArrTree := []int{0}
	// testArrTree := []int{1, 2, 2, 0, 3, 0, 3}
	testArrTree := []int{1, 2, 2, 3, 0, 3, 0}
	// testArrTree := []int{1, 2, 2, 3, 4, 4, 3}
	// testArrTree := []int{5, 1, 4, 0, 0, 3, 6}
	testTree := makeBinaryTree(testArrTree, 0)
	// testTree := &TreeNode{Val: 0}
	// fmt.Println(isValidBST(testTree))
	fmt.Println(isSymmetric2(testTree))
}

// TreeNode struct {
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.0%的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了5.49%的用户
// iteration
func isSymmetric(root *TreeNode) bool {

	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)

	}
	return true
}

// recursiono

// 执行用时：4 ms, 在所有 Go 提交中击败了73.13%的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了35.25%的用户
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true

	}
	return mirror(root, root)
}
func mirror(m, n *TreeNode) bool {

	if m == nil && n == nil {
		return true
	}
	if (m == nil && n != nil) || (m != nil && n == nil) {
		return false
	}
	// if m == nil || n == nil{
	// 	if
	// }
	if m.Val != n.Val {
		return false
	}
	return mirror(m.Left, n.Right) && mirror(m.Right, n.Left)
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
