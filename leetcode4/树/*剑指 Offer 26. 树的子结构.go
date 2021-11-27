/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSubStructure(A *TreeNode, B *TreeNode) bool {
     //IM：  init false
    res:=false
    if A == nil || B == nil {
        return false
    }
    if A.Val == B.Val {
        res = hasSub(A, B)
    }
    if res == false {
        res = isSubStructure(A.Left, B)
    }
    if res == false {
        res = isSubStructure(A.Right, B)
    }
    return res
}
// IM: hasSub
func hasSub(A *TreeNode, B *TreeNode) bool {
    //1
    if B == nil { return true }
    //2
    if A == nil { return false }
    //3
    if A.Val != B.Val { return false}
    //4
    return hasSub(A.Left, B.Left) && hasSub(A.Right, B.Right)
}