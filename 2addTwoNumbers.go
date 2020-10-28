// 2. 两数相加
// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// test := []int{2, 6, 4, 2}
	// Node := makeListNode(test)
	// for i := 0; i < len(test); i++ {
	// 	fmt.Println(Node.Val)
	// 	Node = Node.Next
	// }
	// t1 := []int{1}
	t1 := []int{0}
	// t1 := []int{2, 4, 9}
	// t2 := []int{1, 0}
	t2 := []int{0}
	// t2 := []int{5, 6, 4, 9}
	n1 := makeListNode(t1)
	n2 := makeListNode(t2)
	result := addTwoNumbers(n1, n2)
	for result.Next != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
	fmt.Println(result.Val)

}

func makeListNode(nums []int) *ListNode {
	// var	root *ListNode
	root := &ListNode{Val: nums[0]}
	// if len(nums) <= 1 {
	// 	return root
	// }
	head := root
	for i := 1; i < len(nums); i++ {
		Node := ListNode{Val: nums[i]}
		// Node.Val = nums[i]
		root.Next = &Node
		root = root.Next
	}
	return head
}

// 执行用时：16 ms, 在所有 Go 提交中击败了38.14%的用户
// 内存消耗：4.9 MB, 在所有 Go 提交中击败了45.70%的用户
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode
	var head *ListNode
	t := 0
	for l1 != nil || l2 != nil {
		t1, t2 := 0, 0
		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		}
		t = t + t1 + t2
		if head == nil {
			head = &ListNode{Val: t % 10}
			root = head
		} else {
			root.Next = &ListNode{Val: t % 10}
			root = root.Next
		}
		t = t / 10

	}
	if t != 0 {
		Node := &ListNode{Val: t}
		t = t / 10
		root.Next = Node
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	anchor := &ListNode{}
	head := anchor
	left := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		temp := left + v1 + v2
		val := temp % 10
		left = temp / 10
		head.Next = &ListNode{val, nil}
		head = head.Next
	}
	if left != 0 {
		head.Next = &ListNode{left, nil}
	}
	return anchor.Next
}
