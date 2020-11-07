// 102. 二叉树的层序遍历
// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

// 示例：
// 二叉树：[3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：

// [
//   [3],
//   [9,20],
//   [15,7]
// ]
package main

import "fmt"

func main() {
	// testArrTree := []int{1, 2, 3, 0, 5, 0, 7, 0, 0, 10, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 5, 0, 7, 0, 0, 10, 0, 0, 0, 14, 15}
	// testArrTree := []int{1, 2, 3, 0, 0, 6, 0}
	testArrTree := []int{1, 0, 3}
	// testArrTree := []int{1, 0, 3, 0, 0, 4, 0}
	testTree := makeBinaryTree(testArrTree, 0)
	fmt.Println(levelOrder(testTree))
}

//此方法速度慢，存储大，不好
// func levelOrder(root *TreeNode) [][]int {
// 	var res [][]int
// 	if root == nil {
// 		return res
// 	}
// 	res = son(0, root, res)
// 	// res = res[:len(res)-1]
// 	for i := len(res) - 1; i >= 0; i-- {
// 		if len(res[i]) == 0 {
// 			res = res[:len(res)-1]
// 		}
// 	}
// 	return res

// }
func son(lv int, root *TreeNode, re [][]int) [][]int {

	// var re [][]int

	// if root.Left != nil || root.Right != nil {
	// 	re = append(re, []int{})
	// }
	re = append(re, []int{})
	re[lv] = append(re[lv], root.Val)

	if root.Left != nil {
		re = son(lv+1, root.Left, re)
	}

	if root.Right != nil {
		re = son(lv+1, root.Right, re)
	}
	return re
}

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

// 这个是别人的方法，我来看看
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		l := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				l = append(l, node.Left)
			}
			if node.Right != nil {
				l = append(l, node.Right)
			}
		}
		q = l
	}
	return res
}
