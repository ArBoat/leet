/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		// IM: 注意return
		// 或者明确三种情况， 三种情况不存在包含
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			val := next(root.Right)
			root.Val = val
			root.Right = deleteNode(root.Right, val)
		}
	}
	return root
}

func next(node *TreeNode) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}