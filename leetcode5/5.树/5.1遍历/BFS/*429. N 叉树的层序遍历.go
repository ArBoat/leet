/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	q := []*Node{root}
	for level := 0; len(q) > 0; level++ {
		next := []*Node{}
		res = append(res, []int{})
		for i := 0; i < len(q); i++ {
			node := q[i]
			res[level] = append(res[level], node.Val)
			for _, v := range node.Children {
				next = append(next, v)
			}
		}
		q = next
	}
	return res
}