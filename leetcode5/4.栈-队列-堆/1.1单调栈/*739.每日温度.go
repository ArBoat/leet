func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0) // save index
	res := make([]int, len(temperatures))
	for i, v := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

/*
单调递减栈
1. stack 含义
2. for 而不是 if

*/