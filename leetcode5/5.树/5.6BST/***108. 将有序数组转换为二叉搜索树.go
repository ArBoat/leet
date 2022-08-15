
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return build(nums, 0, len(nums)-1)
}

// 左闭右闭， low, higt 是index
func build(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  build(nums, low, mid-1),
		Right: build(nums, mid+1, high),
	}
}
