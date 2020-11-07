// 141. 环形链表
// 给定一个链表，判断链表中是否有环。

// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

// 如果链表中存在环，则返回 true 。 否则，返回 false 。

// 进阶：

// 你能用 O(1)（即，常量）内存解决此问题吗？

// 示例 1：

// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
// 示例 2：

// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。
// 示例 3：

// 输入：head = [1], pos = -1
// 输出：false
// 解释：链表中没有环。

// 提示：

// 链表中节点的数目范围是 [0, 104]
// -105 <= Node.val <= 105
// pos 为 -1 或者链表中的一个 有效索引 。
package main

import "fmt"

func main() {
	// p1 := ListNode{
	// 	Val: 1,
	// }
	// p1.makeNodeLength(3)
	// tmpN := recursionMakeNode(4)
	// printNode(tmpN)

	// fmt.Printf("result:%t\n",isPalindrome(tmpN))

	// makeCycleNode(5,2)

	// fmt.Printf("site:%d\n",hasCycle(makeCycleNode(5,2)))
	var testNode *ListNode
	// testNode = makeCycleNode(4,2)
	testNode = makeNodeLength2(2)
	printNode(testNode)
	fmt.Printf("hhhhh:%t\n", hasCycle2(testNode))
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

//两种办法，快慢指针相遇有环，或者数组存地址，最后一个地址重复有环
/*
这是快慢指针方法
执行用时 :12 ms, 在所有 Go 提交中击败了73.23%的用户
内存消耗 :3.8 MB, 在所有 Go 提交中击败了71.75%的用户
*/
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}
	var fast *ListNode
	var slow *ListNode
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/*
数组存地址
执行用时 :12 ms, 在所有 Go 提交中击败了71.85%的用户
内存消耗 :4.7 MB, 在所有 Go 提交中击败了16.51%的用户
*/
func hasCycle2(head *ListNode) bool {
	var sitePoint []*ListNode
	for head != nil {
		tPoint := head
		head = head.Next
		sitePoint = append(sitePoint, tPoint)
		tPoint.Next = nil
	}
	long := len(sitePoint)
	for i := 0; i < long-1; i++ {
		if sitePoint[i] == sitePoint[long-1] {
			return true
		}
	}
	return false
}

//我放弃了空间复杂度未0（1）的进阶要求
//打算用一个切片来记录地址
//这里我把题理解错了，只需要返回bool而不是确定位置
func UnhasCycle(head *ListNode) int {
	var sitePoint []*ListNode
	if head.Next == nil {
		return -1
	}
	for head.Next != nil {
		sitePoint = append(sitePoint, head)
		point := head
		head = head.Next
		point.Next = nil
	}
	//因为一些语法的问题，先搁置在这
	//这里需要返回的是head的地址在sitePoint中的位置
	return 0
}

func (NodeA *ListNode) makeNodeLength(n int) {
	if n <= 1 {
		return
	}
	var point *ListNode
	point = NodeA
	for i := 2; i <= n; i++ {
		newNode := ListNode{Val: i}
		for point.Next != nil {
			point = point.Next
		}
		point.Next = &newNode
	}
}
func makeNodeLength2(n int) *ListNode {
	// var head *ListNode
	// head.Val = 1
	var point *ListNode
	head := ListNode{Val: 1}
	// head.Val = 1
	// point := head
	point = &head
	for i := 2; i <= n; i++ {
		tmpNode := ListNode{Val: i}
		point.Next = &tmpNode
		point = point.Next
	}
	return &head
}

//创造一个环形列表方便测试
func makeCycleNode(long int, site int) *ListNode {
	// var head *ListNodo
	var tmpPoint *ListNode
	var Point *ListNode
	if long <= 0 || site <= 0 {
		return nil
	}
	head := ListNode{Val: 1}
	// head.Val = 1
	if long == 1 {
		return &head
	}
	tmpPoint = &head
	for i := 2; i <= long; i++ {
		newNode := ListNode{Val: i}
		tmpPoint.Next = &newNode
		tmpPoint = tmpPoint.Next
	}
	Point = &head
	for i := 0; i < site; i++ {

		Point = Point.Next
	}

	tmpPoint.Next = Point

	return &head
}
