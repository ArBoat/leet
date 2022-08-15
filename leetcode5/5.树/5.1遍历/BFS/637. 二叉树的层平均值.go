/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := []float64{}
	q := []*TreeNode{root}
	for level := 0; len(q) > 0; level++ {
		next := []*TreeNode{}
		sum := 0
		for i := 0; i < len(q); i++ {
			node := q[i]
			sum += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(q)))
		q = next
	}
	return res
}