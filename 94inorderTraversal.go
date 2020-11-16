// 94. 二叉树的中序遍历
// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。

// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]

// 示例 2：
// 输入：root = []
// 输出：[]
// 示例 3：

// 输入：root = [1]
// 输出：[1]

// 示例 4：
// 输入：root = [1,2]
// 输出：[2,1]

// 示例 5：
// 输入：root = [1,null,2]
// 输出：[1,2]

// 提示：
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100

// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
package main

import "fmt"

func main() {

	//
	// test := [][]byte{
	// 	{'1'},
	// 	{'2', '3'},
	// 	{'4', '5', '6', '7'},
	// }
	testArrTree := []int{1, 2, 2, 3, 0, 3, 0}
	testTree := makeBinaryTree(testArrTree, 0)
	fmt.Println(inorderTraversal(testTree))
}

// func inorderTraversal(root *TreeNode) (res []int) {
// 	var inorder func(node *TreeNode)
// 	inorder = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		inorder(node.Left)
// 		res = append(res, node.Val)
// 		inorder(node.Right)
// 	}
// 	inorder(root)
// 	return
// }

func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
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
