/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums, 0, len(nums))
}

func construct(nums []int, l, r int) *TreeNode {
	//terminal
	if l == r {
		return nil
	}
	index := findMaxIndex(nums, l, r)
	root := &TreeNode{
		Val:   nums[index],
		Left:  construct(nums, l, index),
		Right: construct(nums, index+1, r),
	}
	return root
}

/* [) */
func findMaxIndex(nums []int, l, r int) int {
	max := l
	for i := l; i < r; i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}
	return max
}

// BST 范围搜索和传递
