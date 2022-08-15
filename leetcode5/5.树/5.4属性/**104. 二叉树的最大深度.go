/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	q := []*TreeNode{root}
	for level = 0; len(q) > 0; level++ {
		next := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			node := q[i]
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return level
}