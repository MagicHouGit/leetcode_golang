package main

//TreeNode leftf right
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func levelOrder(root *TreeNode) []int {

// }
func sortedArrayToBST(nums []int) *TreeNode {

	return sor(nums, 0, len(nums)-1)
}
func sor(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sor(nums, left, mid-1)

	root.Right = sor(nums, mid+1, right)

	return root
}
