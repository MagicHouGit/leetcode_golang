// 160. 相交链表
// 编写一个程序，找到两个单链表相交的起始节点。

// 如下面的两个链表：

// 在节点 c1 开始相交。

// 示例 1：

// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
// 输出：Reference of the node with value = 8
// 输入解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

// 示例 2：

// 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// 输出：Reference of the node with value = 2
// 输入解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

// 示例 3：

// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null
// 输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
// 解释：这两个链表不相交，因此返回 null。

// 注意：

// 如果两个链表没有交点，返回 null.
// 在返回结果后，两个链表仍须保持原有的结构。
// 可假定整个链表结构中没有循环。
// 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

package main

import "fmt"

func main() {
	t1 := []int{4, 1, 8, 4, 5}
	t2 := []int{5, 0, 1, 8, 4, 5}
	root1, root2 := mf(t1, t2)

	fmt.Println(root1.Val)
	fmt.Println(root2.Val)
	res := getIntersectionNode(root1, root2)
	fmt.Println(res.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 执行用时：60 ms, 在所有 Go 提交中击败了7.74%的用户
// 内存消耗：7.2 MB, 在所有 Go 提交中击败了49.11%的用户
// func getIntersectionNode(headA, headB *ListNode) *ListNode {

// 	rootA := headA
// 	rootB := headB
// 	for i := 0; i < 3; {
// 		if headA == headB {
// 			return headA
// 		} else {
// 			headA = headA.Next
// 			headB = headB.Next
// 		}
// 		if headA == nil {
// 			headA = rootB
// 			i++
// 		}
// 		if headB == nil {
// 			headB = rootA
// 			i++
// 		}
// 	}
// 	return nil
// }

func mf(nums1 []int, nums2 []int) (head1, head2 *ListNode) {

	l1 := len(nums1)
	l2 := len(nums2)
	longEqual := 0
	for i := 0; ; i++ {
		if nums1[l1-1-i] != nums2[l2-1-i] {
			longEqual = i
			break
		}
	}
	var equalHead *ListNode
	// var head1 *ListNode

	// var root1 *ListNode
	root1 := &ListNode{Val: nums1[0]}
	head1 = root1
	for i := 1; i < l1; i++ {
		temp := &ListNode{Val: nums1[i]}
		root1.Next = temp
		root1 = root1.Next
		if i == l1-longEqual {
			equalHead = root1
		}

	}
	// var head2 *ListNode
	root2 := &ListNode{Val: nums2[0]}
	head2 = root2
	for i := 1; i < l2-longEqual; i++ {
		temp := &ListNode{Val: nums2[i]}
		root2.Next = temp
		root2 = root2.Next
	}
	root2.Next = equalHead

	return head1, head2
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := 0
	tmpA := headA
	for tmpA != nil {
		lengthA++
		tmpA = tmpA.Next
	}
	lengthB := 0
	tmpB := headB
	for tmpB != nil {
		lengthB++
		tmpB = tmpB.Next
	}
	if lengthA > lengthB {
		for lengthA != lengthB {
			headA = headA.Next
			lengthA--
		}
	} else {
		for lengthA != lengthB {
			headB = headB.Next
			lengthB--
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}
	return nil
}
