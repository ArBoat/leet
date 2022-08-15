/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// IM: 三个条件
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

//迭代
func isSymmetric(root *TreeNode) bool {
	p, q := root, root
	stack := []*TreeNode{p, q}
	for len(stack) > 0 {
		p, q = stack[0], stack[1]
		stack = stack[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		stack = append(stack, p.Left)
		stack = append(stack, q.Right)

		stack = append(stack, p.Right)
		stack = append(stack, q.Left)
	}
	return true
}