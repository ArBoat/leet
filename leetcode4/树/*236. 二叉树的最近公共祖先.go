/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil { return nil }
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}
	// 返回的既是公共节点， 也是一个点的本身
	l:=lowestCommonAncestor(root.Left, p, q)  
	r:=lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}