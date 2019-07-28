/*
 */ //
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
