/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    return construct(nums, 0, len(nums))
}

func construct(nums []int, l int, r int) *TreeNode {
    if l == r { return nil }
    // 结构体指针初始化
    root:=&TreeNode{}
    index:= maxIndex(nums, l, r)
    root:=&TreeNode{
        Val:nums[index],
        Left: construct(nums,l,index),
        Right: construct(nums, index+1,r),
    }
    return root
}
// 最大index
func maxIndex(nums []int, l int, r int) int {
    max:= l
    for i:=l;i<r;i++{
        if nums[i]>nums[max] {
            max = i
        }
    }
    return max
}