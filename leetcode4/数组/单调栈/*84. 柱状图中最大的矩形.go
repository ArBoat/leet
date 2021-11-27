func largestRectangleArea(heights []int) int {
    if len(heights) == 0 { return 0 }
    res:=0
    stack, newHeights:= []int{}, []int{0}
    newHeights = append(newHeights, heights...)
    newHeights = append(newHeights, 0)
    for i, v := range newHeights {
        for len(stack)>0 && newHeights[stack[len(stack)-1]] > v {
            left:= stack[len(stack)-2]
            h := stack[len(stack)-1]
            res = max(res,  (i - left - 1)* newHeights[h])
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return res
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*

O(n)/O(n)
单调递增栈
前后各加一个0
*/