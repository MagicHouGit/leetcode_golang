// 剑指 Offer 24. 反转链表
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL

// 限制：

// 0 <= 节点个数 <= 5000
package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5}
	Node := makeNode(test)

	// for Node != nil {
	// 	fmt.Println(Node.Val)
	// 	Node = Node.Next
	// }
	reNode := reverseList(Node)
	for reNode != nil {
		fmt.Println(reNode.Val)
		reNode = reNode.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	var temp *ListNode
// 	for head != nil {
// 		temp = &ListNode{Val: head.Val, Next: temp}
// 		head = head.Next
// 	}
// 	return temp
// }

func makeNode(nums []int) *ListNode {

	root := &ListNode{Val: nums[0]}
	head := root
	for i := 1; i < len(nums); i++ {
		temp := &ListNode{Val: nums[i]}
		root.Next = temp
		root = root.Next
	}
	return head

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	pre := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	//把那个成环的开头那两个，打断
	head.Next = nil

	return pre
}
