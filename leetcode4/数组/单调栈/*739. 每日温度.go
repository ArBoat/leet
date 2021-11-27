func dailyTemperatures(temperatures []int) []int {
    if len(temperatures) == 0 { return nil}
    if len(temperatures) == 1 { return []int{0} }
    // stack 存下标, res先开好空间
    stack, res :=make([]int,0), make([]int, len(temperatures))
    for i, v:= range temperatures {
        for len(stack)>0 && temperatures[stack[len(stack)-1]] < v {
            res[stack[len(stack)-1]] = i - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return res
}

// O(n)/O(n) 辅助栈
//单调递减栈