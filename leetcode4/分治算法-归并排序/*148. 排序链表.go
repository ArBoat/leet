/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
    nums:=[]int{}
    p:=head
    for p != nil {
        nums = append(nums, p.Val)
        p = p.Next
    }
    qucikSelect(nums, 0, len(nums)-1)
    p=head
    for _, v:= range nums {
        p.Val = v
        p = p.Next
    }
    return head
}

func qucikSelect(nums []int, left, right int) {
    if left < right {
        p := partition(nums, left, right)
        qucikSelect(nums, left, p-1)
        qucikSelect(nums, p+1, right)
    }
}

func partition(nums []int, left, right int) int {
    m:= left + (right -left)/2
    nums[m], nums[right] = nums[right], nums[m]
    i, j := left, left
    for ;i<right;i++{
        if nums[i]<nums[right]{
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
    nums[j], nums[right] = nums[right], nums[j]
    return j
}