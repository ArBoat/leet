/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    res := false
    if head == nil || head.Next == nil {
        return true
    }
    slow, fast := head, head
    var pre *ListNode
    for {
        if fast == nil {
            res = compareTwoList(pre, slow)
            break
        } else if fast.Next == nil {
            res = compareTwoList(pre, slow.Next)
            break
        } else {
            fast = fast.Next.Next
            //revert
            //slow, pre, slow.Next = slow.Next, slow, pre
            pre, slow, slow.Next = slow, slow.Next, pre
        }
    }
    //restore
    for pre != nil{
        pre, slow, pre.Next = pre.Next, pre, slow
    }
    return res

}
func compareTwoList(pre , s *ListNode) bool {
    // pre,s一定非空, 且长度相等
    for pre != nil {
        // 判断值
        if pre.Val != s.Val {
            return false
        }
        pre = pre.Next
        s = s.Next
    }
    return true
}

/*
O(n)/O(1)
注意细节
快慢指针， 慢指针同时反转， 快指针结束时判断， 然后复原
*/