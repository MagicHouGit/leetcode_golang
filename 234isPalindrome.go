package main

import "fmt"

func main() {
	//创造一个回文联表
	tmpN := recursionMakeNode(4)
	printNode(tmpN)
	fmt.Printf("result:%t\n", isPalindrome(tmpN))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//遍历输出链表
func printNode(NodeP *ListNode) {
	for NodeP != nil {
		fmt.Printf("Val:%d\n", NodeP.Val)
		NodeP = NodeP.Next
	}
}

//反转链表
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

//执行用时 :28 ms, 在所有 Go 提交中击败了46.04%的用户
//内存消耗 :6.5 MB, 在所有 Go 提交中击败了33.61%的用户

//时间复杂度O(n)有点难啊，
//先试一下快慢指针的方法
//需要调用反转链表reverseList()
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var fast_point *ListNode
	var slow_point *ListNode
	fast_point = head
	slow_point = head

	for fast_point.Next != nil && fast_point.Next.Next != nil {
		slow_point = slow_point.Next
		fast_point = fast_point.Next.Next
	}
	// slow_point = slow_point.Next
	slow_point = reverseList(slow_point.Next)
	for slow_point != nil {
		if slow_point.Val == head.Val {
			slow_point = slow_point.Next
			head = head.Next
		} else {
			return false
		}
	}
	return true
}

//n指的是i递增到的最大值而不是链表的长度
//这里没有用递归
//创造一个回文链表方便测试
func recursionMakeNode(n int) *ListNode {

	var point *ListNode
	if n < 1 {
		return nil
	}

	head := ListNode{Val: 1}
	if n == 1 {
		return &head
	}
	point = &head
	for i := 2; i <= n; i++ {
		newNode := ListNode{Val: i}
		point.Next = &newNode
		point = point.Next
	}
	for j := n; j > 0; j-- {
		newNode := ListNode{Val: j}
		point.Next = &newNode
		point = point.Next
	}
	return &head
}

// 执行用时：16 ms, 在所有 Go 提交中击败了60.75%的用户
// 内存消耗：6.9 MB, 在所有 Go 提交中击败了27.42%的用户
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	var arr []int
	for head.Next != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	arr = append(arr, head.Val)
	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] == arr[l-i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func makeNode(nums []int) *ListNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	root := &ListNode{Val: nums[0], Next: nil}
	res := root
	for i := 1; i < l; i++ {
		temp := &ListNode{Val: nums[i], Next: nil}
		root.Next = temp
		root = root.Next
	}
	return res
}
