func largestRectangleArea(heights []int) int {
	res := 0
	stack := []int{} // index
	newHeights := []int{0}
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, 0)
	for i, v := range newHeights {
		for len(stack) > 0 && newHeights[stack[len(stack)-1]] > v {
			l := stack[len(stack)-2]
			h := newHeights[stack[len(stack)-1]]
			res = max(res, (i-l-1)*h)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
要找每个柱子左右两边第一个小于该柱子的柱子，所以从栈头（元素从栈头弹出）到栈底的顺序应该是从大到小的顺序
O(n)/O(n)
单调非递减栈
前后各加一个0
*/