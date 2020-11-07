// 559. N叉树的最大深度
// 给定一个 N 叉树，找到其最大深度。

// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

// 例如，给定一个 3叉树 :

// 我们应返回其最大深度，3。

// 说明:

// 树的深度不会超过 1000。
// 树的节点总不会超过 5000。
package main

import "fmt"

func main() {

	testNode5 := Node{Val: 5}
	testNode6 := Node{Val: 6}
	testNode3 := Node{Val: 3, Children: []*Node{&testNode5, &testNode6}}
	testNode2 := Node{Val: 2}
	testNode4 := Node{Val: 4}
	testNode := Node{Val: 1, Children: []*Node{&testNode2, &testNode3, &testNode4}}
	fmt.Println(maxDepth(&testNode))

	// eNode1 := Node{Val: 1}
	// eNode2 := Node{Val: 2, Children: []*Node{&eNode1}}
	// eNode3 := Node{Val: 3, Children: []*Node{&eNode2}}
	// fmt.Println(maxDepth(&eNode3))
}

// Node n叉树
type Node struct {
	Val      int
	Children []*Node
}

// 执行用时：4 ms, 在所有 Go 提交中击败了9.72%的用户
// 内存消耗：3.3 MB, 在所有 Go 提交中击败了58.33%的用户
func maxDepth(root *Node) int {
	var res int = 0
	var lp int = 0

	if root != nil {
		res++
	} else {
		return res
	}
	if root.Children != nil {
		for _, v := range root.Children {
			t := maxDepth(v)
			if t > lp {
				lp = t
			}

		}

	}
	res = lp + res
	return res
}
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Children != nil {
				queue = append(queue, curr.Children...)
			}
		}
		depth++
	}
	return depth
}

// func makeBinaryTree(x []int, n int) *TreeNode {

// 	//n初始化为0，用于递归中确定位置
// 	//计算层数 ，比如第一层，是0.第二层 是 1，2。第三层是 3，4，5，6.
// 	/*
// 		通过数组长度先确定层数，
// 		利用层数来循环
// 		需要记录上一层的节点，方便下一层使用（有点占内存，有没有更好的办法？）

// 	*/
// 	if x == nil || x[n] == 0 {
// 		return nil
// 	}
// 	l := len(x)
// 	// var newTreeNode *TreeNode
// 	// newTreeNode.Val = x[n]
// 	// newTreeNode.Val = 1
// 	newTreeNode := TreeNode{Val: x[n]}

// 	if 2*n+1 < l {
// 		newTreeNode.Left = makeBinaryTree(x, 2*n+1)
// 	}
// 	if 2*n+2 < l {
// 		newTreeNode.Right = makeBinaryTree(x, 2*n+2)
// 	}
// 	return &newTreeNode
// }
