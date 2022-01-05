/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	carry := 0
	// 不要初始化
	var res *ListNode
	for len(stack1) != 0 || len(stack2) != 0 || carry != 0 {
		a, b := 0, 0
		if len(stack1) != 0 {
			a = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			b = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		cur := a + b + carry
		carry = cur / 10
		cur %= 10
		curNode := &ListNode{cur, res}
		res = curNode
	}
	return res
}

/*
栈实现，注意res声明时为nil
*/

/*
翻转相加再翻转（看题目需求）
*/