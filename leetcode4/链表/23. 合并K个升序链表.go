/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
    // IM : 判断条件
    if l == r {
        return lists[l]
    }
    if l>r {
        return nil
    }
    m:=l+(r-l)/2
    return mergeTwoLists(merge(lists,l,m),merge(lists,m+1,r))
}

 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy:=&ListNode{0,nil}
    p,a,b:=dummy,l1,l2
    for a != nil && b != nil {
        if a.Val < b.Val {
            p.Next = a
            a=a.Next
        } else {
            p.Next = b
            b=b.Next
        }
        p = p.Next
   }
   if a != nil {
       p.Next = a
   }
   if b != nil {
       p.Next = b
   }
   return dummy.Next
}
/*
O(knlogk)/O(1)
用分治法加速
*/