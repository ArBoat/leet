/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }
    cur := head
    count:=0
    for cur != nil {
        count++
        cur = cur.Next
    }
    res:=make([]int, count)
    //IM: do not forget to reset
    cur = head
    for cur != nil {
        res[count-1] = cur.Val
        count--
        cur = cur.Next
    }
    return res
}