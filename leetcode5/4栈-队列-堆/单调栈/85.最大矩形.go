func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	res := 0
	for _, row := range matrix {
		for j, column := range row {
			if column == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		res = max(res, largest(heights))
	}
	return res
}

func largest(heights []int) int {
	res := 0
	stack := []int{}
	newHeights := []int{0}
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, 0)
	for i, v := range newHeights {
		// im : 单调栈for
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
单调栈  for for for
累计heights数组， 1： +=1， 0： = 0
O(mn) /O(mn)
84题pulus
*/