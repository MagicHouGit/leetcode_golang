// 404. 左叶子之和
// 计算给定二叉树的所有左叶子之和。

// 示例：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

package main

import "fmt"

func main() {

	// testArrTree := []int{1, 2, 2, 3, 0, 3, 0}
	// testArrTree := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	testArrTree := []int{1, 2, 3, 4, 5, 0, 0}
	testTree := makeBinaryTree(testArrTree, 0)
	fmt.Println(sumOfLeftLeaves(testTree))
	fmt.Println(bfs(testTree))
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
	newTreeNode := TreeNode{Val: x[n]}

	if 2*n+1 < l {
		newTreeNode.Left = makeBinaryTree(x, 2*n+1)
	}
	if 2*n+2 < l {
		newTreeNode.Right = makeBinaryTree(x, 2*n+2)
	}
	return &newTreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var res int
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
			if root != nil && root.Right == nil && root.Left == nil {
				res = res + root.Val
			}
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

// func isLeafNode(node *TreeNode) bool {
// 	return node.Left == nil && node.Right == nil
// }

// func dfs(node *TreeNode) (ans int) {
// 	if node.Left != nil {
// 		if isLeafNode(node.Left) {
// 			ans += node.Left.Val
// 		} else {
// 			ans += dfs(node.Left)
// 		}
// 	}
// 	if node.Right != nil && !isLeafNode(node.Right) {
// 		ans += dfs(node.Right)
// 	}
// 	return
// }

// func sumOfLeftLeaves(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return dfs(root)
// }

func bfs(p *TreeNode) []int {
	res := make([]int, 0)
	if p == nil {
		return res
	}
	queue := []*TreeNode{p}
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res = append(res, queue[0].Val)
			queue = queue[1:]
		}
	}
	return res
}
