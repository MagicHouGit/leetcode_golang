//删除链表的第n个节点
package main

func main() {
	p1 := ListNode{
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
	removeNthFromEnd(&p1, 1)
	// p1.Next = p1.Next.Next
	// fmt.Printf("Val:%d\n", p1.Val)
	// fmt.Printf("Val:%d\n", p1.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 执行用时 :4 ms, 在所有 Go 提交中击败了64.58%的用户
// 内存消耗 :2.3 MB, 在所有 Go 提交中击败了6.22%的用户
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//链表只有一个元素
	if head.Next == nil {
		head = head.Next
		return head
	}
	m := map[int]*ListNode{}
	nu := 1
	var tmpNode *ListNode
	var lNode *ListNode
	// var tNode *ListNode
	tmpNode = head
	m[nu] = head
	for tmpNode.Next != nil {
		nu++
		m[nu] = tmpNode.Next
		tmpNode = tmpNode.Next
	}
	//如果删除的是头结点
	if nu == n {
		// head.Val = 32
		// head = head.Next
		// head.Val = 23
		head.Val = head.Next.Val
		head.Next = head.Next.Next
		return head
	}
	// m[nu-n+1]
	// tmpNode = &m[nu-n]
	// tmpNode = &m(nu - n + 1)
	lNode = m[nu-n].Next
	lNode.Next = m[nu-n+1].Next
	m[nu-n].Next = m[nu-n+1].Next

	return head

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

//leetcode上别人的方法学习一下

// 删除链表的倒数第n个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return nil
	}

	// 采用先后指针，first先走n个结点，next再开始走，直到first到达尾部
	first, next := head, head

	// first向前走n个结点
	for n > 0 && first != nil {
		first = first.Next
		n--
	}
	// 如果链表长度不足n，则返回
	if n > 0 {
		return head
	}
	// 如果链表长度恰好为n，则删除头结点
	if first == nil {
		return head.Next
	}

	// 让first多走一步，使得next指向该删除的结点的前一个，便于删除
	first = first.Next
	for first != nil {
		first = first.Next
		next = next.Next
	}
	// 删除next此时的后一个结点
	next.Next = next.Next.Next

	return head
}

// func addNodeN(n int){
//     if n<1{
// 		return
// 	}
// 	if n == 1{
//         node := ListNode{
// 			Val:n,
// 		}
// 		return
// 	}
// 	for i:=n;i<=0;i--{
// 		node := ListNode{
// 			Val:n,
// 			Next:,
// 		}
// 	}

// }
