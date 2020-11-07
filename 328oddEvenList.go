// 328. 奇偶链表
// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

// 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

// 示例 1:

// 输入: 1->2->3->4->5->NULL
// 输出: 1->3->5->2->4->NULL
// 示例 2:

// 输入: 2->1->3->5->6->4->7->NULL
// 输出: 2->3->6->7->1->5->4->NULL
// 说明:

// 应当保持奇数节点和偶数节点的相对顺序。
// 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7}
	Node := manufactureNode(test)
	// for Node != nil {
	// 	fmt.Println(Node.Val)
	// 	Node = Node.Next
	// }
	result := oddEvenList(Node)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.2 MB, 在所有 Go 提交中击败了33.33%的用户
func oddEvenList(head *ListNode) *ListNode {
	var even *ListNode
	var odd *ListNode
	odd = head
	even = head.Next
	evneHead := even
	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next

	}
	odd.Next = evneHead
	return head

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func manufactureNode(nums []int) *ListNode {
	var head *ListNode
	root := &ListNode{Val: nums[0]}
	head = root
	for i := 1; i < len(nums); i++ {
		Node := &ListNode{Val: nums[i]}
		root.Next = Node
		root = root.Next
	}
	return head
}
