func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	start, end := 0, 0
	res, sum := n+1, 0
	for ; end < n; end++ {
		sum += nums[end]
		for sum >= target {
			res = min(res, end-start+1)
			sum -= nums[start]
			start++
		}
	}
	//IM： 不存在target的情况
	if res == n+1 {
		return 0
	} else {
		return res
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}