/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    p:=head
    for p != nil {
        temp:= &Node{p.Val, nil, nil}
        temp.Next = p.Next
        //1
        p.Next = temp
        p = p.Next.Next
    }
    p=head
    for p!=nil {
        //2
        if p.Random != nil {
            p.Next.Random = p.Random.Next
        }
        p = p.Next.Next
    }
    p=head
    res, q:=head.Next, head.Next
    // IM: 一定先p，再q
    for q.Next != nil {
        p.Next = p.Next.Next
        q.Next = q.Next.Next
        p = p.Next
        q = q.Next
    }
    // 注意3
    p.Next = nil
    return res
}