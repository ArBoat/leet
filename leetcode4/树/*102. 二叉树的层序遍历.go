/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    res:=[][]int{}
    var dfs func(root *TreeNode, level int)
    dfs =func(root *TreeNode, level int){
        if root == nil { return }
        if level == len(res){
            res = append(res, []int{})
        }
        res[level] = append(res[level], root.Val)
        //IM : level+1
        dfs(root.Left, level+1)
        dfs(root.Right, level+1)
    }
    // do not forget to call
    dfs(root, 0)
    return res
}