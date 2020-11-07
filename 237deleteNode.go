/*
 */ //
//  237. 删除链表中的节点
//  请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。

//  现有一个链表 -- head = [4,5,1,9]，它可以表示为:

//  示例 1：

//  输入：head = [4,5,1,9], node = 5
//  输出：[4,1,9]
//  解释：给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//  示例 2：

//  输入：head = [4,5,1,9], node = 1
//  输出：[4,5,9]
//  解释：给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

//  提示：

//  链表至少包含两个节点。
//  链表中所有节点的值都是唯一的。
//  给定的节点为非末尾节点并且一定是链表中的一个有效节点。
//  不要从你的函数中返回任何结果。
package main

import (
	"fmt"
)

// type Post struct {
// 	body        string
// 	publishDate int64 // Unix timestamp
// 	next        *Post // link to the next Post
// }
// type Feed struct {
// 	length int // we'll use it later
// 	start  *Post
// }

// func (f *Feed) Append(newPost *Post) {
// 	if f.length == 0 {
// 		f.start = newPost
// 	} else {
// 		currentPost := f.start
// 		for currentPost.next != nil {
// 			currentPost = currentPost.next
// 		}
// 		currentPost.next = newPost
// 	}
// 	f.length++
// }

// func main() {
// 	f := &Feed{}
// 	p1 := Post{
// 		body: "Lorem ipsum",
// 	}
// 	f.Append(&p1)

// 	fmt.Printf("Length: %v\n", f.length)
// 	fmt.Printf("First: %v\n", f.start)

// 	p2 := Post{
// 		body: "Dolor sit amet",
// 	}
// 	f.Append(&p2)

// 	fmt.Printf("Length: %v\n", f.length)
// 	fmt.Printf("First: %v\n", f.start)
// 	fmt.Printf("Second: %v\n", f.start.next)
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func delete(node *ListNode) {
	// if head.Val == node.Val{
	// 	head = head.Next
	// }else{
	// 	tempNode:= head.Next
	// 	if tempNode.Val == node.Val{
	// 		head.Next == tempNode.Next
	// 	}else{
	// 		tempNode =
	// 	}
	// }
	var tempNode *ListNode
	tempNode = node.Next
	node.Val = tempNode.Val
	node.Next = tempNode.Next

}
func (NodeA *ListNode) AddNode(node *ListNode) {

	if NodeA.Next == nil {
		NodeA.Next = node
	} else {
		tmpNext := NodeA.Next
		for tmpNext == nil {
			tmpNext = NodeA.Next
		}
		// NodeA.Next = node
		tmpNext.Next = node
	}

}
func main() {
	p1 := ListNode{
		Val: 1,
	}
	p2 := ListNode{
		Val: 2,
	}
	p3 := ListNode{
		Val: 3,
	}
	p4 := ListNode{
		Val: 4,
	}
	p1.AddNode(&p2)
	// p1.AddNode(&p3)
	fmt.Printf("val:%d\n", p1.Next.Val)
	// fmt.Printf("val:%d\n", p2.Next.Val)
	p1.AddNode(&p3)
	fmt.Printf("val:%d\n", p1.Next.Val)
	p1.AddNode(&p4)
	fmt.Printf("val:%d\n", p1.Next.Val)

}
