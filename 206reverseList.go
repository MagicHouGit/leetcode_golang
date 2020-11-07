// 206反转链表
// 反转一个单链表。

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 执行用时 :4 ms, 在所有 Go 提交中击败了83.58%的用户
// 内存消耗 :3.1 MB, 在所有 Go 提交中击败了5.27%的用户

package main

import "fmt"

func main() {
	p1 := ListNode{
		Val: 1,
	}

	p1.makeNodeLength(5)
	tmp1 := &p1
	for tmp1.Next != nil {
		fmt.Printf("Val:%d\n", tmp1.Val)
		tmp1 = tmp1.Next
	}
	fmt.Printf("Val:%d\n", tmp1.Val)
	var tmp2 *ListNode
	tmp2 = reverseList(&p1)
	for tmp2.Next != nil {
		fmt.Printf("Val:%d\n", tmp2.Val)
		tmp2 = tmp2.Next
	}
	fmt.Printf("Val:%d\n", tmp2.Val)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var tmpNode *ListNode
	if head == nil || head.Next == nil {
		return head
	} else {
		tmpNode = reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return tmpNode
	}
}

func deleteNode(node *ListNode) {
	var tempNode *ListNode
	tempNode = node.Next
	node.Val = tempNode.Val
	node.Next = tempNode.Next
}

func (NodeA *ListNode) addNode(node *ListNode) {
	if NodeA.Next == nil {
		NodeA.Next = node
	} else {
		tmpNode := NodeA.Next
		for tmpNode.Next != nil {
			tmpNode = tmpNode.Next
		}
		tmpNode.Next = node
	}

}

//制造一个制定长的的链表，
func (NodeA *ListNode) makeNodeLength(n int) {
	if n <= 1 {
		return
	}
	var point *ListNode
	// var newNode *ListNode
	point = NodeA
	for i := 2; i <= n; i++ {
		// newNode.Val = n
		newNode := ListNode{Val: i}
		for point.Next != nil {

			point = point.Next
		}
		point.Next = &newNode
	}

}
