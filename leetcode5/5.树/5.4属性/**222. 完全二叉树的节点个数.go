/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := countHeight(root.Left)
	r := countHeight(root.Right)
	// IM : 位运算
	if l == r {
		return (1 << l) + countNodes(root.Right)
	} else {
		return (1 << r) + countNodes(root.Left)
	}
}

func countHeight(root *TreeNode) int {
	res := 0
	for root != nil {
		res++
		root = root.Left
	}
	return res
}

//直接计数， 不好
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}