/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		next := []*TreeNode{}
		res = append(res, max(q))
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
	return res
}

func max(q []*TreeNode) int {
	max := math.MinInt32
	for _, v := range q {
		if v.Val > max {
			max = v.Val
		}
	}
	return max
}