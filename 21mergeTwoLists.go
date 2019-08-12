//21合并两个有效链表


//执行用时 :4 ms, 在所有 Go 提交中击败了81.10%的用户
//内存消耗 :2.6 MB, 在所有 Go 提交中击败了60.05%的用户


package main

import "fmt"

func main() {
	p1 := ListNode{
		Val: 1,
	}
	p2 := ListNode{
		Val: 1,
	}
	// p2 := ListNode{
	// 	Val: 2,
	// }
	// p3 := ListNode{
	// 	Val: 3,
	// }
	// p4 := ListNode{
	// 	Val: 4,
	// }
	// p5 := ListNode{
	// 	Val: 5,
	// }
	// p1.addNode(&p2)
	// fmt.Printf("Val:%d\n", p1.Next.Val)
	// p1.addNode(&p3)
	// fmt.Printf("Val:%d\n", p2.Next.Val)
	// p1.addNode(&p4)
	// fmt.Printf("Val:%d\n", p3.Next.Val)
	// p1.addNode(&p5)
	// deleteNode(&p3)
	// fmt.Printf("Val:%d\n", p2.Next.Val)
	// removeNthFromEnd(&p1, 1)
	// p1.Next = p1.Next.Next
	// fmt.Printf("Val:%d\n", p1.Val)
	// fmt.Printf("Val:%d\n", p1.Next.Val)

	p1.makeNodeLength(3)

	tmp1 := &p1
	for tmp1.Next != nil {
		fmt.Printf("Val:%d\n",tmp1.Val)
		tmp1 = tmp1.Next
	}
	fmt.Printf("Val:%d\n",tmp1.Val)
	
	p2.makeNodeLength2(3)

	// var tmp2 *ListNode
	// tmp2 = reverseList(&p1)
	tmp2 := &p2
	for tmp2.Next != nil {
		fmt.Printf("Val:%d\n",tmp2.Val)
		tmp2 = tmp2.Next
	}
	fmt.Printf("Val:%d\n",tmp2.Val)



	tmp3 := mergeTwoLists(&p1,&p2)
	for tmp3.Next != nil {
		fmt.Printf("Val:%d\n",tmp3.Val)
		tmp3 = tmp3.Next
	}
	fmt.Printf("Val:%d\n",tmp3.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var tmpNode *ListNode
	if head == nil || head.Next == nil{
		return head
	}else{
		tmpNode = reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return tmpNode
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l1.Val <= l2.Val{
		
		l1.Next = mergeTwoLists(l1.Next,l2)
		// l1.Next = l2
		return l1
	}else{

		l2.Next = mergeTwoLists(l1,l2.Next)
		// l2.Next = l1
		return l2
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
func (NodeA *ListNode) makeNodeLength(n int){
    if n<=1 {
		return 
	}
	var point *ListNode
	// var newNode *ListNode
	point = NodeA
	for i:=2;i<=n;i++{
		// newNode.Val = n       
		newNode := ListNode{Val:i}
		for point.Next != nil{

			point= point.Next
		}
        point.Next = &newNode
	}
}
func (NodeA *ListNode) makeNodeLength2(n int){
	if n<=1 {
		return 
	}
	var point *ListNode
	point = NodeA
	for i:=2;i<=n;i++{
		newNode := ListNode{Val:i*2}
		for point.Next != nil{
			point = point.Next
		}
		point.Next = &newNode
	}
}