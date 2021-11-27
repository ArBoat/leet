/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 //DFS
 func zigzagLevelOrder(root *TreeNode) [][]int {
    res:=[][]int{}
    var dfs func(root *TreeNode, level int)
    dfs = func(root *TreeNode, level int) {
        if root == nil { return }
        if level == len(res) {
            res = append(res, []int{})
        }
        if level%2 == 0 {
            res[level] = append(res[level], root.Val)
        } else {
            res[level] = append([]int{root.Val}, res[level]...)
        }
        dfs(root.Left, level+1)
        dfs(root.Right, level+1)
    }
    dfs(root, 0)
    return res
}

//BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil { return nil}
    res:=[][]int{}
    q:=[]*TreeNode{root}
    for i:=0;len(q)>0;i++{
        res = append(res, []int{})
        next :=[]*TreeNode{}
        for j:=0;j<len(q);j++{
            node:=q[j]
            if i%2==0{
                res[i] = append(res[i],node.Val)
            } else {
                res[i] = append([]int{node.Val}, res[i]...)
            }
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