func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if num <= i { //IM: check
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}
	return dp[target]
}

/*
有序， 其实是排列
dp[i] count of  target is j
dp[i] = dp[i] + dp[i-nums[j]]
dp[0] = 1
*/