/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//DFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//IM: 考虑左右子树分别是否为空的情况，共4种
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	level := 0
	for ; len(q) > 0; level++ {
		next := []*TreeNode{}
		for _, v := range q {
			if v.Left == nil && v.Right == nil {
				return level + 1
			}
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		q = next
	}
	return level + 1
}