// 103. 二叉树的锯齿形层次遍历
// 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

// 例如：
// 给定二叉树 [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回锯齿形层次遍历如下：

// [
//   [3],
//   [20,9],
//   [15,7]
// ]

package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7}
	testTree := makeBinaryTree(test, 0)
	fmt.Println(bfs(testTree))
	fmt.Println(zigzagLevelOrder(testTree))
}

// func cursion
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

func bfs(root *TreeNode) []int {
	res := make([]int, 0)

	queue := []*TreeNode{root}
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

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了71.72%的用户
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	LeftStart := true
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		es := make([]int, length)
		for i := 0; i < length; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if LeftStart {
				es[i] = queue[0].Val
			} else {
				es[length-i-1] = queue[0].Val
			}
			// length--
			queue = queue[1:]
		}
		res = append(res, es)
		LeftStart = !LeftStart
	}
	return res

}
