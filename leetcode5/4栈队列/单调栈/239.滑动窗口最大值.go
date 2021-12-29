func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}
	stack := []int{}
	push := func(i int) {
		// 注意用for
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	// res的值都nums的元素
	res := []int{nums[stack[0]]}
	for i := k; i < n; i++ {
		push(i)
		for stack[0] <= i-k {
			stack = stack[1:]
		}
		res = append(res, nums[stack[0]])
	}
	return res
}