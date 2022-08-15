/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	var traversal func(root *TreeNode, path string)
	traversal = func(root *TreeNode, path string) {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			res = append(res, path)
			return
		}
		if root.Left != nil {
			traversal(root.Left, path+"->")
		}
		if root.Right != nil {
			traversal(root.Right, path+"->")
		}
	}
	traversal(root, "")
	return res
}
