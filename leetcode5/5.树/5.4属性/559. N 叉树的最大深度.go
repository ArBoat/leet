/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	res := 0
	for _, v := range root.Children {
		res = max(res, maxDepth(v))
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}