func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 { //改变原数组,累加
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

/*
贪心
O(n)/O(1)
*/

/*
dp:
dp[i] = max(dp[i - 1] + nums[i], nums[i])
*/