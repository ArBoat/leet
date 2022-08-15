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
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	m := l + (r-l)/2
	return mergeTwoLists(merge(lists, l, m), merge(lists, m+1, r))
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

/*
O(knlogk)/O(1)
用二分/分治法加速
*/