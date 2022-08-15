/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//BFS
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		next := []*TreeNode{}
		//im: type
		res = q[0].Val
		for _, v := range q {
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		q = next
	}
	return res
}
